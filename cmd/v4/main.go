package main

import (
	"context"
	"encoding/json"
	"log"
	"robanohashi/internal/config"
	"robanohashi/internal/model"
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

	subjectKeys, err := db.Client().Keys(context.Background(), "subject:*").Result()

	if err != nil {
		log.Fatal(err)
	}

	deleteCount := 0

	vSet := make(map[string]bool)

	for _, key := range subjectKeys {

		raw, err := db.JSONGet(context.Background(), key)

		if err != nil {
			log.Fatal(err)
		}

		var rawJson map[string]interface{}

		err = json.Unmarshal([]byte(raw.(string)), &rawJson)

		if err != nil {
			log.Fatal(err)
		}

		if rawJson["object"] != "vocabulary" {
			continue
		}

		var v model.Vocabulary

		err = json.Unmarshal([]byte(raw.(string)), &v)

		if err != nil {
			log.Fatal(err)
		}

		if _, isset := vSet[v.Characters]; !isset {

			vSet[v.Characters] = true
			continue
		}

		for _, kanjiId := range v.ComponentSubjectIds {

			k, err := db.GetKanji(context.Background(), kanjiId)

			if err != nil {
				log.Fatal(err)
			}

			k.AmalgamationSubjectIds = remove(k.AmalgamationSubjectIds, v.ID)

			err = db.JSONSet(keys.Subject(kanjiId), "$", k)

			if err != nil {
				log.Fatal(err)
			}

		}

		err = db.Client().Del(context.Background(), key).Err()

		if err != nil {
			log.Fatal(err)
		}

		deleteCount++

		if deleteCount%1000 == 0 {
			log.Printf("Deleted %d subjects sofar", deleteCount)
		}
	}

	log.Printf("Deleted %d subjects", deleteCount)
}

func remove[T comparable](l []T, item T) []T {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}
