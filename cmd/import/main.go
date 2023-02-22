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
			InsertKanji(context.Background(), rdb, kanji)
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

func InsertKanji(ctx context.Context, rdb *redis.Client, kanji Subject[Kanji]) {
	key := fmt.Sprintf("kanji:%d", kanji.ID)
	err := rdb.HSet(ctx, key,
		"id", kanji.ID,
		"object", string(kanji.Object),
		"characters", kanji.Data.Characters,
		"slug", kanji.Data.Slug,
		"reading_mnemonic", createReadingMnemonic(&kanji),
	).Err()

	if err != nil {
		log.Fatal(err)
	}

	if len(kanji.Data.AmalgamationSubjectIds) > 0 {
		key = fmt.Sprintf("kanji:%d:amalgamation_ids", kanji.ID)
		_, err = rdb.SAdd(ctx, key, kanji.Data.AmalgamationSubjectIds...).Result()
		if err != nil {
			log.Fatal(err)
		}
	}

	key = fmt.Sprintf("kanji:%d:meanings", kanji.ID)
	_, err = rdb.SAdd(ctx, key, createKanjiMeanings(&kanji)...).Result()
	if err != nil {
		log.Fatal(err)
	}

	key = fmt.Sprintf("kanji:%d:readings", kanji.ID)
	_, err = rdb.SAdd(ctx, key, createKanjiReadings(&kanji)...).Result()
	if err != nil {
		log.Fatal(err)
	}

	key = fmt.Sprintf("kanji:%d:component_subject_ids", kanji.ID)
	_, err = rdb.SAdd(ctx, key, kanji.Data.ComponentSubjectIds...).Result()
	if err != nil {
		log.Fatal(err)
	}

	if len(kanji.Data.VisuallySimilarSubjectIds) > 0 {
		key = fmt.Sprintf("kanji:%d:visually_similar_subject_ids", kanji.ID)
		_, err = rdb.SAdd(ctx, key, kanji.Data.VisuallySimilarSubjectIds...).Result()
		if err != nil {
			log.Fatal(err)
		}
	}

	key = fmt.Sprintf("kanji:%d:meaning_mnemonics", kanji.ID)

	_, err = rdb.SAdd(ctx, key, createMeaningMnemonic(&kanji)).Result()
	if err != nil {
		log.Fatal(err)
	}
}

type Meaning struct {
	Meaning string `json:"meaning"`
	Primary bool   `json:"primary"`
}

func (m Meaning) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

type KanjiReading struct {
	Reading string `json:"reading"`
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
}

func (r KanjiReading) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func createMeaningMnemonic(kanji *Subject[Kanji]) string {
	meaningMnemonic := kanji.Data.MeaningMnemonic
	if kanji.Data.MeaningHint != "" {
		meaningMnemonic = meaningMnemonic + " " + kanji.Data.MeaningHint
	}
	return meaningMnemonic
}

func createKanjiReadings(kanji *Subject[Kanji]) []any {
	readings := make([]any, 0)
	for _, reading := range kanji.Data.Readings {
		readings = append(readings, KanjiReading{
			Reading: reading.Reading,
			Primary: reading.Primary,
			Type:    reading.Type,
		})
	}

	return readings
}

func createReadingMnemonic(kanji *Subject[Kanji]) string {
	readingMnemonic := kanji.Data.ReadingMnemonic
	if kanji.Data.ReadingHint != "" {
		readingMnemonic = readingMnemonic + " " + kanji.Data.ReadingHint
	}
	return readingMnemonic
}

func createKanjiMeanings(kanji *Subject[Kanji]) []any {
	meanings := make([]any, 0)
	for _, meaning := range kanji.Data.Meanings {
		meanings = append(meanings, Meaning{
			Meaning: meaning.Meaning,
			Primary: meaning.Primary,
		})
	}
	for _, aux_meaning := range kanji.Data.AuxiliaryMeanings {
		meanings = append(meanings, Meaning{
			Meaning: aux_meaning.Meaning,
			Primary: false,
		})
	}

	return meanings
}
