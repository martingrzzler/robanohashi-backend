package persist

import (
	"context"
	"robanohashi/persist/keys"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

type DB struct {
	rdb         *redis.Client
	jsonHandler *rejson.Handler
}

func (db *DB) JSONHandler() *rejson.Handler {
	return db.jsonHandler
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

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		return nil, err
	}

	handler := rejson.NewReJSONHandler()
	handler.SetGoRedisClient(rdb)

	return &DB{rdb: rdb, jsonHandler: handler}, nil
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
	err = db.rdb.Do(context.Background(),
		"FT.CREATE",
		keys.SearchIndex(),
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

func (db *DB) Close() {
	db.rdb.Close()
}
