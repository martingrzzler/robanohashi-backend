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
	"robanohashi/persist/keys"
)

type Meaning struct {
	Meaning string `json:"meaning"`
	Primary bool   `json:"primary"`
}

type Reading struct {
	Reading string `json:"reading"`
	Type    string `json:"type"`
}

type KanjiData struct {
	Characters string    `json:"characters"`
	Object     string    `json:"object"`
	Meanings   []Meaning `json:"meanings"`
	Readings   []Reading `json:"readings"`
	Slug       string    `json:"slug"`
	Radicals   []string  `json:"radicals"`
}

func main() {
	nextId := 9166

	cfg := config.New()
	db, err := persist.Connect(cfg.RedisURL, cfg.RedisPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	f, err := os.Open("kanjis.jsonl")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		var data KanjiData

		err := json.Unmarshal([]byte(line), &data)

		if err != nil {
			log.Fatal(err)
		}

		res, err := db.SearchSubjects(context.Background(), data.Characters)

		if err != nil {
			log.Fatal(err)
		}

		if res.TotalCount > 0 {
			continue
		}

		kanji := model.Kanji{
			ID:                        nextId,
			Object:                    "kanji",
			Slug:                      data.Slug,
			Characters:                data.Characters,
			VisuallySimilarSubjectIds: []int{},
			AmalgamationSubjectIds:    []int{},
			Meanings:                  copyMeanings(data.Meanings),
			Readings:                  copyReadings(data.Readings),
			ComponentSubjectIds:       []int{},
		}

		for _, radical := range data.Radicals {
			res, err := db.SearchSubjects(context.Background(), radical)

			if err != nil {
				log.Fatal(err)
			}

			for _, subject := range res.Items {
				if subject.Object == "radical" && subject.Characters == radical {
					kanji.ComponentSubjectIds = append(kanji.ComponentSubjectIds, subject.ID)
					break
				}

			}
		}

		err = db.JSONSet(keys.Subject(kanji.ID), "$", kanji)
		if err != nil {
			log.Fatal(err)
		}

		nextId += 1
	}
}

func copyMeanings(meanings []Meaning) []model.Meaning {
	var result []model.Meaning

	for _, meaning := range meanings {
		result = append(result, model.Meaning{
			Meaning: meaning.Meaning,
			Primary: meaning.Primary,
		})
	}

	return result
}

func copyReadings(readings []Reading) []model.KanjiReading {
	var result []model.KanjiReading

	for _, reading := range readings {
		result = append(result, model.KanjiReading{
			Reading: reading.Reading,
			Type:    reading.Type,
			Primary: false,
		})
	}

	return result
}
