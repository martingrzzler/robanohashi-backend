package main

import (
	"context"
	"fmt"
	"log"
	"robanohashi/internal/config"
	"robanohashi/internal/model"
	"robanohashi/persist"
	"strings"
)

type MnemonicData struct {
	KanjiID    int    `json:"kanji_id"`
	Characters string `json:"characters"`
	Mnemonic   string `json:"mnemonic"`
}

func main() {
	cfg := config.New()

	db, err := persist.Connect(cfg.RedisURL, cfg.RedisPassword)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	kanjiMap := make(map[string]bool)

	keys, err := db.Client().Keys(context.Background(), "meaning_mnemonic:*").Result()

	if err != nil {
		log.Fatal(err)
	}

	for _, key := range keys {

		id := strings.Split(key, ":")[1]

		mm, err := db.GetMeaningMnemonic(context.Background(), id)

		if err != nil {
			log.Fatal(err)
		}

		if mm.UserID == string(model.NonHumanUserIDAIGenerated) {

			if kanjiMap[mm.SubjectID] {
				panic(fmt.Sprintf("duplicate kanji id %s", mm.SubjectID))
			} else {
				kanjiMap[mm.SubjectID] = true
			}

		}
	}
	fmt.Println(len(kanjiMap))
}

// func main() {

// 	insertedCount := 0
// 	cfg := config.New()

// 	db, err := persist.Connect(cfg.RedisURL, cfg.RedisPassword)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()

// 	f, err := os.Open("mnemonics_last.jsonl")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	scanner := bufio.NewScanner(f)

// 	for scanner.Scan() {

// 		data := scanner.Text()

// 		var m MnemonicData

// 		err := json.Unmarshal([]byte(data), &m)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		kanji, err := db.GetKanji(context.Background(), m.KanjiID)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		mm := model.MeaningMnemonic{
// 			ID:          uuid.New().String(),
// 			Text:        m.Mnemonic,
// 			VotingCount: 0,
// 			SubjectID:   strconv.Itoa(kanji.ID),
// 			UserID:      string(model.NonHumanUserIDAIGenerated),
// 			CreatedAt:   time.Now().Unix(),
// 			UpdatedAt:   time.Now().Unix(),
// 		}

// 		err = db.JSONSet(keys.MeaningMnemonic(mm.ID), "$", mm)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		insertedCount++
// 	}

// 	fmt.Printf("Inserted %d mnemonics\n", insertedCount)
// }
