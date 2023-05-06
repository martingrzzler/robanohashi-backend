package main

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"
	"robanohashi/internal/config"
	"robanohashi/internal/model"
	"robanohashi/persist"
)

func main() {

	cfg := config.New()

	db, err := persist.Connect(cfg.RedisURL, cfg.RedisPassword)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	filename := "kanjis.jsonl"

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(f)

	keys, err := db.Client().Keys(context.Background(), "kanji:*").Result()

	if err != nil {
		log.Fatal(err)
	}

	for _, key := range keys {

		raw, err := db.JSONGet(context.Background(), key)

		if err != nil {
			log.Fatal(err)
		}

		var kanji model.Kanji

		err = json.Unmarshal([]byte(raw.(string)), &kanji)

		if err != nil {
			log.Fatal(err)
		}

		if kanji.Source != model.SourceWaniKani {
			continue
		}

		kres, err := db.GetKanjiResolved(context.Background(), &kanji)

		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := json.Marshal(kres)

		if err != nil {
			log.Fatal(err)
		}

		_, err = writer.WriteString(string(jsonData) + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	err = writer.Flush()

	if err != nil {
		log.Fatal(err)
	}
}
