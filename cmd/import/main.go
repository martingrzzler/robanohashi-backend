package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func main() {

	rdb := Connect()

	f, err := os.Open("subjects.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var data map[string]interface{}

		json.Unmarshal([]byte(line), &data)

		switch Object(data["object"].(string)) {
		case ObjectKanji:
			kanji := Subject[Kanji]{}
			err = json.Unmarshal([]byte(line), &kanji)
			if err != nil {
				log.Fatal(err)
			}
			key := fmt.Sprintf("kanji:%d", kanji.ID)
			_, err = rdb.HSet(context.Background(), key, "characters", kanji.Data.Characters).Result()
			if err != nil {
				log.Fatal(err)
			}

		case ObjectRadical:
			radical := Subject[Radical]{}
			err = json.Unmarshal([]byte(line), &radical)
			if err != nil {
				log.Fatal(err)
			}
			key := fmt.Sprintf("radical:%d", radical.ID)
			_, err = rdb.HSet(context.Background(), key, "characters", radical.Data.Characters).Result()
			if err != nil {
				log.Fatal(err)
			}
		case ObjectVocabulary:
			vocabulary := Subject[Vocabulary]{}
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
