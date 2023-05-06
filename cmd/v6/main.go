package main

import (
	"context"
	"encoding/json"
	"log"
	"robanohashi/internal/config"
	"robanohashi/persist"
	"robanohashi/persist/keys"
)

func main() {

	cfg := config.New()

	db, err := persist.Connect(cfg.RedisURL, cfg.RedisPassword)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	subjectKeys, err := db.Client().Keys(context.Background(), "vocabulary:*").Result()

	if err != nil {
		log.Fatal(err)
	}

	for _, key := range subjectKeys {

		raw, err := db.JSONGet(context.Background(), key)

		if err != nil {
			log.Fatal(err)
		}

		var vocabulary map[string]interface{}

		err = json.Unmarshal([]byte(raw.(string)), &vocabulary)

		if err != nil {
			log.Fatal(err)
		}

		if vocabulary["source"] == "wanikani" {
			vocabulary["source"] = 0
		} else {
			vocabulary["source"] = 1
		}

		err = db.JSONSet(keys.Vocabulary(int(vocabulary["id"].(float64))), "$", vocabulary)

		if err != nil {
			log.Fatal(err)
		}
	}

}
