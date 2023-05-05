package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"robanohashi/internal/config"
	"robanohashi/internal/model"
	"robanohashi/persist"
	"robanohashi/persist/keys"
)

type Reading struct {
	Reading string `json:"reading"`
	Romaji  string `json:"romaji"`
	Primary bool   `json:"primary"`
}

type VocabData struct {
	Characters string   `json:"characters"`
	Meanings   []string `json:"meanings"`
	Reading    Reading  `json:"reading"`
	Kanjis     []string `json:"kanjis"`
}

func main() {
	nextId := 17072

	cfg := config.New()
	db, err := persist.Connect(cfg.RedisURL, cfg.RedisPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	f, err := os.Open("vocabs.jsonl")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var data VocabData

		err := json.Unmarshal([]byte(line), &data)

		if err != nil {
			log.Fatal(err)
		}

		v := model.Vocabulary{
			ID:                  nextId,
			Object:              "vocabulary",
			Slug:                data.Characters,
			Characters:          data.Characters,
			ComponentSubjectIds: []int{},
			Meanings:            createMeanings(data.Meanings),
			Readings:            []model.VocabularyReading{copyReading(data.Reading)},
			ContextSentences:    []model.ContextSentence{},
		}

		for _, kanji := range data.Kanjis {
			res, err := db.SearchSubjects(context.Background(), kanji)

			if err != nil {
				log.Fatal(err)
			}

			for _, subject := range res.Items {
				if subject.Object == "kanji" && subject.Characters == kanji {
					v.ComponentSubjectIds = append(v.ComponentSubjectIds, subject.ID)
					break
				}

			}

		}

		if len(v.ComponentSubjectIds) != len(data.Kanjis) {
			panic(fmt.Sprintf("Missing kanji for %s", data.Characters))
		}

		err = db.JSONSet(keys.Subject(v.ID), "$", v)
		if err != nil {
			log.Fatal(err)
		}

		nextId += 1

		if (nextId-17072)%10000 == 0 {
			fmt.Println((float32(nextId) - 17071.0) / 297822.0)
		}

	}
}

func copyReading(r Reading) model.VocabularyReading {
	return model.VocabularyReading{
		Reading: r.Reading,
		Romaji:  r.Romaji,
		Primary: r.Primary,
	}
}

func createMeanings(ms []string) []model.Meaning {
	meanings := make([]model.Meaning, len(ms))

	for i, m := range ms {
		meanings[i] = model.Meaning{
			Meaning: m,
			Primary: i == 0,
		}
	}

	return meanings
}
