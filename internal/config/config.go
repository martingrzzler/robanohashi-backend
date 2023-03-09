package config

import "os"

type Config struct {
	RedisURL      string
	RedisPassword string
}

func NewConfig() Config {
	return Config{
		RedisURL:      os.Getenv("REDIS_URL"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}
}
