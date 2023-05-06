package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"robanohashi/internal/dto"
	"robanohashi/internal/model"
	"robanohashi/persist/keys"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var upvoteMnemonicScript = redis.NewScript(`
local uid = ARGV[1]
local mnemonicKey = KEYS[1]
local mnemonicUpvotesKey = KEYS[2]
local mnemonicDownvotesKey = KEYS[3]

if redis.call("SISMEMBER", mnemonicUpvotesKey, uid) == 1 then
	redis.call("SREM", mnemonicUpvotesKey, uid)
	redis.call("JSON.NUMINCRBY", mnemonicKey, "$.voting_count", -1)
	return {ok = "removed upvote"}
end

if redis.call("SISMEMBER", mnemonicDownvotesKey, uid) == 1 then
	redis.call("SMOVE", mnemonicDownvotesKey, mnemonicUpvotesKey, uid)
	redis.call("JSON.NUMINCRBY", mnemonicKey, "$.voting_count", 2)
	return {ok = "switched from downvote to upvote"}
end

redis.call("SADD", mnemonicUpvotesKey, uid)
redis.call("JSON.NUMINCRBY", mnemonicKey, "$.voting_count", 1)

return {ok = "upvoted"}
`)

var downvoteMnemonicScript = redis.NewScript(`
local uid = ARGV[1]
local mnemonicKey = KEYS[1]
local mnemonicUpvotesKey = KEYS[2]
local mnemonicDownvotesKey = KEYS[3]

if redis.call("SISMEMBER", mnemonicDownvotesKey, uid) == 1 then
	redis.call("SREM", mnemonicDownvotesKey, uid)
	redis.call("JSON.NUMINCRBY", mnemonicKey, "$.voting_count", 1)
	return {ok = "removed downvote"}
end

if redis.call("SISMEMBER", mnemonicUpvotesKey , uid) == 1 then
	redis.call("SMOVE", mnemonicUpvotesKey, mnemonicDownvotesKey, uid)
	redis.call("JSON.NUMINCRBY", mnemonicKey, "$.voting_count", -2)
	return {ok = "switched from upvote to downvote"}
end

redis.call("SADD", mnemonicDownvotesKey, uid)
redis.call("JSON.NUMINCRBY", mnemonicKey, "$.voting_count", -1)

return {ok = "downvoted"}
`)

func (db *DB) UpvoteMeaningMnemonic(ctx context.Context, mid string, uid string) (string, error) {
	keys := []string{keys.MeaningMnemonic(mid), keys.MeaningMnemonicUpVoters(mid), keys.MeaningMnemonicDownVoters(mid)}
	argv := []interface{}{uid}

	status, err := upvoteMnemonicScript.Run(ctx, db.rdb, keys, argv).Result()

	if err != nil {
		return "", err
	}

	return status.(string), err
}

func (db *DB) DownvoteMeaningMnemonic(ctx context.Context, mid string, uid string) (string, error) {
	keys := []string{keys.MeaningMnemonic(mid), keys.MeaningMnemonicUpVoters(mid), keys.MeaningMnemonicDownVoters(mid)}
	argv := []interface{}{uid}

	status, err := downvoteMnemonicScript.Run(ctx, db.rdb, keys, argv).Result()

	if err != nil {
		return "", err
	}

	return status.(string), err
}

func (db *DB) ToggleFavoriteMeaningMnemonic(ctx context.Context, mid string, uid string) (string, error) {
	return db.toggleSetValue(ctx, keys.MeaningMnemonicFavorites(uid), mid)
}

func (db *DB) ResolveMeaningMnemonics(ctx context.Context, uid string, mnemonics []model.MeaningMnemonic) ([]dto.MeaningMnemonicWithUserInfo, error) {
	pipe := db.rdb.Pipeline()

	for _, mnemonic := range mnemonics {
		pipe.SIsMember(ctx, keys.MeaningMnemonicUpVoters(mnemonic.ID), uid)
		pipe.SIsMember(ctx, keys.MeaningMnemonicDownVoters(mnemonic.ID), uid)
		pipe.SIsMember(ctx, keys.MeaningMnemonicFavorites(uid), mnemonic.ID)
	}

	res, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	resolved := make([]dto.MeaningMnemonicWithUserInfo, len(mnemonics))

	for i, mnemonic := range mnemonics {
		resolved[i] = dto.MeaningMnemonicWithUserInfo{
			ID:          mnemonic.ID,
			Text:        mnemonic.Text,
			VotingCount: mnemonic.VotingCount,
			UserID:      mnemonic.UserID,
			CreatedAt:   mnemonic.CreatedAt,
			UpdatedAt:   mnemonic.UpdatedAt,
			Upvoted:     res[i*3].(*redis.BoolCmd).Val(),
			Downvoted:   res[i*3+1].(*redis.BoolCmd).Val(),
			Favorite:    res[i*3+2].(*redis.BoolCmd).Val(),
			Me:          mnemonic.UserID == uid,
		}

		sid, _ := strconv.Atoi(mnemonic.SubjectID)

		s, err := db.resolveMnemonicSubject(ctx, sid)
		if err != nil {
			return nil, err
		}
		resolved[i].Subject = s
	}

	return resolved, nil
}

func (db *DB) resolveMnemonicSubject(ctx context.Context, sid int) (dto.MnemonicSubject, error) {
	// identify what type of subject it is
	res, err := db.rdb.Exists(ctx, keys.Kanji(sid)).Result()

	if err != nil {
		return dto.Kanji{}, fmt.Errorf("failed to check if subject is kanji: %w", err)
	}

	// is kanji
	if res == 1 {
		k, err := db.GetKanji(ctx, sid)

		if err != nil {
			return dto.Kanji{}, fmt.Errorf("failed to get kanji: %w", err)
		}

		return db.GetKanjiResolved(ctx, k)
	}

	res, err = db.rdb.Exists(ctx, keys.Vocabulary(sid)).Result()

	if err != nil {
		return dto.Vocabulary{}, fmt.Errorf("failed to check whether vocabulary exists: %w", err)
	}
	// not a kanji nor a vocabulary
	if res == 0 {
		return dto.Vocabulary{}, fmt.Errorf("subject with id %d does not exist", sid)
	}

	v, err := db.GetVocabulary(ctx, sid)

	if err != nil {
		return dto.Vocabulary{}, fmt.Errorf("failed to get vocabulary: %w", err)
	}

	return db.GetVocabularyResolved(ctx, v)
}

func (db *DB) GetMeaningMnemonic(ctx context.Context, id string) (*model.MeaningMnemonic, error) {
	data, err := db.JSONGet(ctx, keys.MeaningMnemonic(id))

	if err != nil {
		return nil, err
	}
	mm := &model.MeaningMnemonic{}

	err = json.Unmarshal([]byte(data.(string)), mm)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal meaning mnemonic: %w", err)
	}

	return mm, nil
}

func (db *DB) UpdateMeaningMnemonic(ctx context.Context, umm dto.UpdateMeaningMnemonic) error {
	tx := db.rdb.TxPipeline()

	tx.Do(ctx, "JSON.SET", keys.MeaningMnemonic(umm.ID), "$.text", fmt.Sprintf("\"%s\"", umm.Text))
	tx.Do(ctx, "JSON.SET", keys.MeaningMnemonic(umm.ID), "$.updated_at", time.Now().Unix())

	_, err := tx.Exec(ctx)

	if err != nil {
		return fmt.Errorf("failed to execute transaction: %w", err)
	}

	return nil
}

func (db *DB) DeleteMeaningMnemonic(ctx context.Context, id string) error {
	err := db.rdb.Do(ctx, "JSON.DEL", keys.MeaningMnemonic(id)).Err()

	if err != nil {
		return fmt.Errorf("failed to delete meaning mnemonic: %w", err)
	}

	return nil
}

func (db *DB) GetFavoriteMeaningMnemonics(ctx context.Context, uid string) ([]model.MeaningMnemonic, error) {
	mids, err := db.rdb.SMembers(ctx, keys.MeaningMnemonicFavorites(uid)).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to get favorite meaning mnemonics: %w", err)
	}

	pipe := db.rdb.Pipeline()
	mnemonicsCmds := make([]*redis.Cmd, len(mids))

	for i, id := range mids {
		mnemonicsCmds[i] = pipe.Do(ctx, "JSON.GET", keys.MeaningMnemonic(id))
	}

	_, err = pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	mnemonics := make([]model.MeaningMnemonic, len(mnemonicsCmds))

	for i, cmd := range mnemonicsCmds {
		m := model.MeaningMnemonic{}
		err = json.Unmarshal([]byte(cmd.Val().(string)), &m)

		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal meaning mnemonic: %w", err)
		}

		mnemonics[i] = m
	}

	return mnemonics, nil
}

func (db *DB) getMeaningMnemonicsBy(ctx context.Context, query string) (*dto.List[model.MeaningMnemonic], error) {
	res, err := db.rdb.Do(context.Background(), "FT.SEARCH", keys.MeaningMnemonicIndex(), query, "SORTBY", "voting_count", "DESC", "LIMIT", "0", "100", "RETURN", "1", "$").Result()

	if err != nil {
		return nil, fmt.Errorf("failed to get meaning mnemonics: %w", err)
	}

	totalCount, items, err := parseFTSearchResult[model.MeaningMnemonic](res)

	if err != nil {
		return nil, fmt.Errorf("failed to parse meaning mnemonics: %w", err)
	}

	return &dto.List[model.MeaningMnemonic]{
		TotalCount: totalCount,
		Items:      items,
	}, nil
}

func (db *DB) GetMeaningMnemonicsByUser(ctx context.Context, uid string) (*dto.List[model.MeaningMnemonic], error) {
	query := fmt.Sprintf("@user_id:{%s}", uid)
	return db.getMeaningMnemonicsBy(ctx, query)
}

func (db *DB) GetMeaningMnemonicsBySubjectID(ctx context.Context, id int) (*dto.List[model.MeaningMnemonic], error) {
	query := fmt.Sprintf("@subject_id:{%d}", id)

	return db.getMeaningMnemonicsBy(ctx, query)
}
