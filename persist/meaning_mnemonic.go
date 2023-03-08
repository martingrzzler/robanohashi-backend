package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"robanohashi/internal/dto"
	"robanohashi/internal/model"
	"robanohashi/persist/keys"
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
	return db.toggleSetValue(ctx, keys.MeaningMnemonicFavorites(mid), uid)
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

	tx.Do(ctx, "JSON.SET", keys.MeaningMnemonic(umm.ID), "$.text", umm.Text)
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
