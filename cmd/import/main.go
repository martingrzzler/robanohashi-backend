package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"robanohashi/cmd/import/wanikani"
	"robanohashi/internal/model"
	"robanohashi/persist"
)

func main() {
	db, err := persist.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	f, err := os.Open("subjects.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0

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
			InsertKanji(context.Background(), db, &kanji)
		case model.ObjectRadical:
			radical := wanikani.Subject[wanikani.Radical]{}
			err = json.Unmarshal([]byte(line), &radical)
			if err != nil {
				log.Fatal(err)
			}
			InsertRadical(context.Background(), db, &radical)
		case model.ObjectVocabulary:
			vocabulary := wanikani.Subject[wanikani.Vocabulary]{}
			err = json.Unmarshal([]byte(line), &vocabulary)
			if err != nil {
				log.Fatal(err)
			}
			InsertVocabulary(context.Background(), db, &vocabulary)
		}

		count++
	}

	fmt.Printf("Successfully imported data %d records", count)
}
