package persist

import (
	"context"
	"encoding/json"
	"fmt"
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

func (db *DB) KeyExists(ctx context.Context, key string) bool {
	res, err := db.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false
	}

	return res == 1
}
