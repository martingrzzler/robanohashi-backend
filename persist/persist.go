package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"robanohashi/internal/dto"
	"robanohashi/internal/model"
	"robanohashi/persist/keys"
	"strings"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	rdb *redis.Client
}

func (db *DB) Client() *redis.Client {
	return db.rdb
}

func Connect() (*DB, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Ping(context.Background()).Err()

	if err != nil {
		return nil, err
	}

	return &DB{rdb: rdb}, nil
}

func (db *DB) CreateIndices() error {
	res, err := db.rdb.Do(context.Background(), "FT._LIST").Result()
	if err != nil {
		return err
	}

	for _, index := range res.([]interface{}) {
		err := db.rdb.Do(context.Background(), "FT.DROPINDEX", index).Err()
		if err != nil {
			return err
		}
	}

	// Create indices
	err = db.createSubjectIndex()
	if err != nil {
		return fmt.Errorf("failed to create subject index: %w", err)
	}

	err = db.createMeaningMnemonicIndex()

	if err != nil {
		return fmt.Errorf("failed to create meaning mnemonic index: %w", err)
	}

	return nil
}

func (db *DB) createMeaningMnemonicIndex() error {
	err := db.rdb.Do(context.Background(),
		"FT.CREATE",
		keys.MeaningMnemonicIndex(),
		"ON", "JSON",
		"PREFIX", "1", "meaning_mnemonic:",
		"SCHEMA",
		"$.subject_id", "AS", "subject_id", "TAG",
		"$.voting_count", "AS", "voting_count", "NUMERIC", "SORTABLE",
	).Err()

	return err
}

func (db *DB) createSubjectIndex() error {
	// Create indices
	err := db.rdb.Do(context.Background(),
		"FT.CREATE",
		keys.SubjectIndex(),
		"ON", "JSON",
		"PREFIX", "3", "kanji:", "radical:", "vocabulary:",
		"SCHEMA",
		"$.characters", "AS", "characters", "TAG",
		"$.meanings.*.meaning", "AS", "meaning", "TEXT",
		"$.readings.*.reading", "AS", "reading", "TAG",
		"$.readings.*.romaji", "AS", "romaji", "TAG",
	).Err()

	return err
}

func (db *DB) JSONSet(key string, value any) error {
	data, err := json.Marshal(value)

	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	return db.rdb.Do(context.Background(), "JSON.SET", key, "$", string(data)).Err()
}

func (db *DB) Close() {
	db.rdb.Close()
}

func (db *DB) SearchSubjects(ctx context.Context, search string) (any, error) {

	query := ""
	if len(strings.Split(search, " ")) > 1 {
		query = fmt.Sprintf("@meaning:(%s*)", search)
	} else {
		query = fmt.Sprintf("((@characters:{%s*}) => { $weight: 2.0 } | (@meaning:(%s*)) | (@reading:{%s*}) | (@romaji:{%s*}))", search, search, search, search)
	}

	return db.rdb.Do(context.Background(), "FT.SEARCH", keys.SubjectIndex(), query, "LIMIT", "0", "20").Result()
}

func (db *DB) GetMeaningMnemonicsBySubjectID(ctx context.Context, id int) (any, error) {
	query := fmt.Sprintf("@subject_id:{%d}", id)

	return db.rdb.Do(context.Background(), "FT.SEARCH", keys.MeaningMnemonicIndex(), query, "SORTBY", "voting_count", "DESC", "LIMIT", "0", "100", "RETURN", "1", "$").Result()
}

func (db *DB) JSONGet(ctx context.Context, key string) (any, error) {
	data, err := db.rdb.Do(context.Background(), "JSON.GET", key).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to get json: %w", err)
	}

	return data, nil
}

type Result struct {
	Data  any
	Error error
}

func (db *DB) GetKanjiResolved(ctx context.Context, kanji *model.Kanji) (*model.ResolvedKanji, error) {
	pipe := db.rdb.Pipeline()
	componentsCmds := make([]*redis.Cmd, len(kanji.ComponentSubjectIds))
	visuallySimCmds := make([]*redis.Cmd, len(kanji.VisuallySimilarSubjectIds))
	amalgamationCmds := make([]*redis.Cmd, len(kanji.AmalgamationSubjectIds))

	for i, id := range kanji.ComponentSubjectIds {
		componentsCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Radical(id))
	}

	for i, id := range kanji.VisuallySimilarSubjectIds {
		visuallySimCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Kanji(id))
	}

	for i, id := range kanji.AmalgamationSubjectIds {
		amalgamationCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Vocabulary(id))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	resolvedKanji := &model.ResolvedKanji{
		ID:                      kanji.ID,
		Slug:                    kanji.Slug,
		Characters:              kanji.Characters,
		Object:                  kanji.Object,
		Meanings:                kanji.Meanings,
		Readings:                kanji.Readings,
		ReadingMnemonic:         kanji.ReadingMnemonic,
		ComponentSubjects:       make([]model.Radical, len(componentsCmds)),
		VisuallySimilarSubjects: make([]model.Kanji, len(visuallySimCmds)),
		AmalgamationSubjects:    make([]model.Vocabulary, len(amalgamationCmds)),
	}

	for i, cmd := range componentsCmds {
		radical := model.Radical{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &radical)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedKanji.ComponentSubjects[i] = radical
	}

	for i, cmd := range visuallySimCmds {
		kanji := model.Kanji{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &kanji)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedKanji.VisuallySimilarSubjects[i] = kanji
	}

	for i, cmd := range amalgamationCmds {
		vocab := model.Vocabulary{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &vocab)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}
		resolvedKanji.AmalgamationSubjects[i] = vocab
	}

	return resolvedKanji, nil
}

func (db *DB) GetKanji(ctx context.Context, id int) (*model.Kanji, error) {
	data, err := db.JSONGet(ctx, keys.Kanji(id))
	if err != nil {
		return nil, err
	}

	kanji := &model.Kanji{}

	err = json.Unmarshal([]byte(data.(string)), kanji)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return kanji, nil
}

func (db *DB) GetRadical(ctx context.Context, id int) (*model.Radical, error) {
	data, err := db.JSONGet(ctx, keys.Radical(id))
	if err != nil {
		return nil, err
	}

	radical := &model.Radical{}

	err = json.Unmarshal([]byte(data.(string)), radical)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return radical, nil
}

func (db *DB) GetRadicalResolved(ctx context.Context, radical *model.Radical) (*model.ResolvedRadical, error) {
	pipe := db.rdb.Pipeline()

	amalgamationCmds := make([]*redis.Cmd, len(radical.AmalgamationSubjectIds))

	for i, id := range radical.AmalgamationSubjectIds {
		amalgamationCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Kanji(id))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	resolvedRadical := &model.ResolvedRadical{
		ID:                   radical.ID,
		Object:               radical.Object,
		Slug:                 radical.Slug,
		Characters:           radical.Characters,
		CharacterImage:       radical.CharacterImage,
		Meanings:             radical.Meanings,
		MeaningMnemonic:      radical.MeaningMnemonic,
		AmalgamationSubjects: make([]model.Kanji, len(amalgamationCmds)),
	}

	for i, cmd := range amalgamationCmds {
		kanji := model.Kanji{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &kanji)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedRadical.AmalgamationSubjects[i] = kanji
	}

	return resolvedRadical, nil
}

func (db *DB) GetVocabulary(ctx context.Context, id int) (*model.Vocabulary, error) {
	data, err := db.JSONGet(ctx, keys.Vocabulary(id))
	if err != nil {
		return nil, err
	}

	vocab := &model.Vocabulary{}

	err = json.Unmarshal([]byte(data.(string)), vocab)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return vocab, nil
}

func (db *DB) GetVocabularyResolved(ctx context.Context, vocab *model.Vocabulary) (*model.ResolvedVocabulary, error) {
	pipe := db.rdb.Pipeline()

	componentCmds := make([]*redis.Cmd, len(vocab.ComponentSubjectIds))

	for i, id := range vocab.ComponentSubjectIds {
		componentCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Kanji(id))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	resolvedVocab := &model.ResolvedVocabulary{
		ID:                vocab.ID,
		Object:            vocab.Object,
		Slug:              vocab.Slug,
		Characters:        vocab.Characters,
		Meanings:          vocab.Meanings,
		ReadingMnemonic:   vocab.ReadingMnemonic,
		Readings:          vocab.Readings,
		ContextSentences:  vocab.ContextSentences,
		ComponentSubjects: make([]model.Kanji, len(componentCmds)),
	}

	for i, cmd := range componentCmds {
		kanji := model.Kanji{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &kanji)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedVocab.ComponentSubjects[i] = kanji
	}

	return resolvedVocab, nil
}

func (db *DB) KeyExists(ctx context.Context, key string) bool {
	res, err := db.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false
	}

	return res == 1
}

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

func (db *DB) ResolveUserVotes(ctx context.Context, uid string, mnemonics []dto.MeaningMnemonic) ([]dto.MeaningMnemonicWithUserInfo, error) {
	pipe := db.rdb.Pipeline()

	for _, mnemonic := range mnemonics {
		pipe.SIsMember(ctx, keys.MeaningMnemonicUpVoters(mnemonic.ID), uid)
		pipe.SIsMember(ctx, keys.MeaningMnemonicDownVoters(mnemonic.ID), uid)
	}

	res, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	votes := make([]dto.MeaningMnemonicWithUserInfo, len(mnemonics))

	for i, mnemonic := range mnemonics {
		votes[i] = dto.MeaningMnemonicWithUserInfo{
			MeaningMnemonic: mnemonic,
			Upvoted:         res[i*2].(*redis.BoolCmd).Val(),
			Downvoted:       res[i*2+1].(*redis.BoolCmd).Val(),
		}
	}

	return votes, nil
}
