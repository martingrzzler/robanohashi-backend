package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"robanohashi/internal/config"
	"robanohashi/internal/dto"
	"robanohashi/internal/model"
	"robanohashi/persist"
	"strconv"
	"strings"
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
	// create csv file
	// loop through all subjects
	// continue if subject is not kanji
	// if subject has no amalgamation_subject_ids add it to the csv file with 0 indicating no amalgamation
	// else add it with the number of amalgamation_subject_ids

	out := "kanjis.jsonl"
	f, err := os.OpenFile(out, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)

	cfg := config.New()
	db, err := persist.Connect(cfg.RedisURL, cfg.RedisPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	keys, err := db.Client().Keys(context.TODO(), "subject:*").Result()

	if err != nil {
		log.Fatal(err)
	}

	for _, key := range keys {

		id, err := strconv.Atoi(strings.Split(key, ":")[1])

		if err != nil {
			log.Fatal(err)
		}

		raw, err := db.JSONGet(context.Background(), key)

		if err != nil {
			log.Fatal(err)
		}

		var data map[string]any

		err = json.Unmarshal([]byte(raw.(string)), &data)

		if err != nil {
			log.Fatal(err)
		}

		if data["object"] != "kanji" {
			continue
		}

		kanji, err := db.GetKanji(context.Background(), id)

		if err != nil {
			log.Fatal(err)
		}

		rsv, err := db.GetKanjiResolved(context.Background(), kanji)

		if err != nil {
			log.Fatal(err)
		}

		jsonLine := struct {
			Characters string   `json:"characters"`
			Vocabulary []string `json:"vocabulary"`
		}{
			Characters: kanji.Characters,
			Vocabulary: getVocabulary(rsv.AmalgamationSubjects),
		}

		jsonBytes, err := json.Marshal(jsonLine)

		if err != nil {
			log.Fatal(err)
		}

		_, err = writer.WriteString(string(jsonBytes) + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing the buffer: %v\n", err)
		return
	}

}

// func main() {
// 	nextId := 9166

// 	cfg := config.New()
// 	db, err := persist.Connect(cfg.RedisURL, cfg.RedisPassword)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	f, err := os.Open("kanjis.jsonl")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		var data KanjiData

// 		err := json.Unmarshal([]byte(line), &data)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		res, err := db.SearchSubjects(context.Background(), data.Characters)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if res.TotalCount > 0 {
// 			continue
// 		}

// 		kanji := model.Kanji{
// 			ID:                        nextId,
// 			Object:                    "kanji",
// 			Slug:                      data.Slug,
// 			Characters:                data.Characters,
// 			VisuallySimilarSubjectIds: []int{},
// 			AmalgamationSubjectIds:    []int{},
// 			Meanings:                  copyMeanings(data.Meanings),
// 			Readings:                  copyReadings(data.Readings),
// 			ComponentSubjectIds:       []int{},
// 		}

// 		for _, radical := range data.Radicals {
// 			res, err := db.SearchSubjects(context.Background(), radical)

// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			for _, subject := range res.Items {
// 				if subject.Object == "radical" && subject.Characters == radical {
// 					kanji.ComponentSubjectIds = append(kanji.ComponentSubjectIds, subject.ID)
// 					break
// 				}

// 			}
// 		}

// 		err = db.JSONSet(keys.Subject(kanji.ID), "$", kanji)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		nextId += 1
// 	}
// }

func getVocabulary(as []dto.SubjectPreview) []string {
	result := make([]string, 0)

	for _, a := range as {
		result = append(result, a.Characters)
	}

	return result
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
