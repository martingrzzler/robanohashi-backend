package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"robanohashi/cmd/import/wanikani"
	"robanohashi/model"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

type Config struct {
	json   *rejson.Handler
	client *redis.Client
}

func main() {

	rdb := Connect()
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

		switch model.Object(data["object"].(string)) {
		case model.ObjectKanji:
			kanji := wanikani.Subject[wanikani.Kanji]{}
			err = json.Unmarshal([]byte(line), &kanji)
			if err != nil {
				log.Fatal(err)
			}
			InsertKanji(context.Background(), cfg, &kanji)
		case model.ObjectRadical:
			radical := wanikani.Subject[wanikani.Radical]{}
			err = json.Unmarshal([]byte(line), &radical)
			if err != nil {
				log.Fatal(err)
			}
			InsertRadical(context.Background(), cfg, &radical)
		case model.ObjectVocabulary:
			vocabulary := wanikani.Subject[wanikani.Vocabulary]{}
			err = json.Unmarshal([]byte(line), &vocabulary)
			if err != nil {
				log.Fatal(err)
			}
			key := fmt.Sprintf("vocabulary:%d", vocabulary.ID)
			_, err = rdb.HSet(context.Background(), key, "characters", vocabulary.Data.Characters).Result()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func Connect() *redis.Client {
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
