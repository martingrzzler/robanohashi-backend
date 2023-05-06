package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"robanohashi/internal/dto"
	robaUtil "robanohashi/internal/util"
	"robanohashi/persist/keys"
	"strings"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	rdb *redis.Client
}

type RawUnmarshaler[T any] interface {
	UnmarshalRaw(data any) (T, error)
}

func parseFTSearchResult[T RawUnmarshaler[T]](result any) (int64, []T, error) {
	items := make([]T, 0)

	for i, item := range result.([]any)[1:] {
		if i%2 == 0 {
			continue
		}

		obj := *new(T)

		parsed, err := obj.UnmarshalRaw(item.([]any)[1])
		if err != nil {
			return 0, nil, err
		}

		items = append(items, parsed)
	}

	totalCount := result.([]interface{})[0].(int64)

	return totalCount, items, nil
}

func (db *DB) Client() *redis.Client {
	return db.rdb
}

func Connect(url string, password string) (*DB, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       0,
	})

	err := rdb.Ping(context.Background()).Err()

	if err != nil {
		return nil, err
	}

	return &DB{rdb: rdb}, nil
}

func (db *DB) JSONSet(key string, path string, value any) error {
	data, err := json.Marshal(value)

	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	return db.rdb.Do(context.Background(), "JSON.SET", key, path, string(data)).Err()
}

func (db *DB) Close() {
	db.rdb.Close()
}

func (db *DB) SearchSubjects(ctx context.Context, search string) (*dto.List[dto.SubjectPreview], error) {
	query := ""
	if len(strings.Split(search, " ")) > 1 {
		query = fmt.Sprintf("@meaning:(%s*)", search)
	} else if robaUtil.SingleKanji(search) {
		query = fmt.Sprintf("@characters:{%s*}", search)
	} else {
		query = fmt.Sprintf("((@characters:{%s*}) => { $weight: 2.0 } | (@meaning:(%s*)) | (@reading:{%s*}) | (@romaji:{%s*}))", search, search, search, search)
	}

	// sort by source, show wanikani first (source 0)
	res, err := db.rdb.Do(context.Background(), "FT.SEARCH", keys.SubjectIndex(), query, "SORTBY", "source", "ASC", "LIMIT", "0", "300", "RETURN", "1", "$").Result()

	if err != nil {
		return nil, fmt.Errorf("failed to search subjects: %w", err)
	}

	totalCount, subjects, err := parseFTSearchResult[dto.SubjectPreview](res)

	if err != nil {
		return nil, fmt.Errorf("failed to parse subjects: %w", err)
	}

	return &dto.List[dto.SubjectPreview]{
		TotalCount: totalCount,
		Items:      subjects,
	}, nil
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

var toggleScript = redis.NewScript(`
local value = ARGV[1]
local key = KEYS[1]

if redis.call("SADD", key, value) == 1 then
	return {ok = "added"}
end

redis.call("SREM", key, value)

return {ok = "removed"}
`)

func (db *DB) toggleSetValue(ctx context.Context, key string, value string) (string, error) {
	keys := []string{key}
	argv := []interface{}{value}

	status, err := toggleScript.Run(ctx, db.rdb, keys, argv).Result()

	if err != nil {
		return "", fmt.Errorf("could not toggle value: %w", err)
	}

	return status.(string), err
}
