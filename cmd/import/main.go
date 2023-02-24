package main

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"robanohashi/cmd/import/wanikani"
	"robanohashi/db"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

type Config struct {
	json   *rejson.Handler
	client *redis.Client
}

func main() {

	rdb := connect()
	defer rdb.Close()

	f, err := os.Open("subjects.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(rdb)

	cfg := Config{
		json:   rh,
		client: rdb,
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var data map[string]interface{}

		json.Unmarshal([]byte(line), &data)

		switch db.Object(data["object"].(string)) {
		case db.ObjectKanji:
			kanji := wanikani.Subject[wanikani.Kanji]{}
			err = json.Unmarshal([]byte(line), &kanji)
			if err != nil {
				log.Fatal(err)
			}
			InsertKanji(context.Background(), cfg, &kanji)
		case db.ObjectRadical:
			radical := wanikani.Subject[wanikani.Radical]{}
			err = json.Unmarshal([]byte(line), &radical)
			if err != nil {
				log.Fatal(err)
			}
			InsertRadical(context.Background(), cfg, &radical)
		case db.ObjectVocabulary:
			vocabulary := wanikani.Subject[wanikani.Vocabulary]{}
			err = json.Unmarshal([]byte(line), &vocabulary)
			if err != nil {
				log.Fatal(err)
			}
			InsertVocabulary(context.Background(), cfg, &vocabulary)
		}
	}
}

func connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	return rdb
}
