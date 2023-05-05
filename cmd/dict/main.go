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

		if !mapKanjis(db, data, &v) {
			continue
		}

		if len(v.ComponentSubjectIds) != len(data.Kanjis) {
			fmt.Println(fmt.Sprintf("Missing kanji for %s", data.Characters))
		}

		err = db.JSONSet(keys.Subject(v.ID), "$", v)
		if err != nil {
			log.Fatal(err)
		}

		nextId += 1

		if (nextId-17072)%1000 == 0 {
			fmt.Println((float32(nextId) - 17071.0) / 297822.0)
		}

	}
}

func mapKanjis(db *persist.DB, d VocabData, v *model.Vocabulary) bool {
	for _, kanji := range d.Kanjis {
		k, err := db.GetKanjiByCharacters(context.Background(), kanji)

		if err != nil {
			fmt.Println("Kanji not found: " + kanji)
			fmt.Println("Vocab: " + d.Characters)

			return false

		}

		v.ComponentSubjectIds = append(v.ComponentSubjectIds, k.ID)

		k.AmalgamationSubjectIds = append(k.AmalgamationSubjectIds, v.ID)

		err = db.JSONSet(keys.Subject(k.ID), "$", k)

		if err != nil {
			return false
		}
	}

	return true
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
