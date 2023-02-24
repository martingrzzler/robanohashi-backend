package db

import (
	"github.com/go-redis/redis/v8"
)

func Connect() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
