package main

import (
	"context"
	"log"
	"robanohashi/cmd/import/wanikani"
	"robanohashi/db"
	"robanohashi/db/keys"
	"strconv"
	"time"
)

func InsertKanji(ctx context.Context, cfg Config, wkKanji *wanikani.Subject[wanikani.Kanji]) {
	err := cfg.client.Incr(ctx, keys.MeaningMnemonicIds()).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := cfg.client.Get(ctx, keys.MeaningMnemonicIds()).Result()
	if err != nil {
		log.Fatal(err)
	}

	id, _ := strconv.Atoi(val)

	meaningMnemonic := db.MeaningMnemonic{
		ID:        id,
		Text:      createKanjiMeaningMnemonic(wkKanji),
		CreatedAt: strconv.FormatInt(time.Now().Unix(), 10),
		UpdatedAt: strconv.FormatInt(time.Now().Unix(), 10),
	}

	_, err = cfg.json.JSONSet(keys.MeaningMnemonic(id), "$", meaningMnemonic)
	if err != nil {
		log.Fatal(err)
	}

	kanji := db.Kanji{
		ID:                        wkKanji.ID,
		Object:                    wkKanji.Object,
		Characters:                wkKanji.Data.Characters,
		Slug:                      wkKanji.Data.Slug,
		ReadingMnemonic:           createReadingMnemonic(wkKanji),
		AmalgamationSubjectIds:    wkKanji.Data.AmalgamationSubjectIds,
		Meanings:                  createKanjiMeanings(wkKanji),
		Readings:                  createKanjiReadings(wkKanji),
		ComponentSubjectIds:       wkKanji.Data.ComponentSubjectIds,
		VisuallySimilarSubjectIds: wkKanji.Data.VisuallySimilarSubjectIds,
		MeaningMnemonicIds:        []int{id},
	}

	_, err = cfg.json.JSONSet(keys.Kanji(wkKanji.ID), "$", kanji)
	if err != nil {
		log.Fatal(err)
	}
}

func createKanjiMeaningMnemonic(kanji *wanikani.Subject[wanikani.Kanji]) string {
	meaningMnemonic := kanji.Data.MeaningMnemonic
	if kanji.Data.MeaningHint != "" {
		meaningMnemonic = meaningMnemonic + " " + kanji.Data.MeaningHint
	}
	return meaningMnemonic
}

func createKanjiReadings(kanji *wanikani.Subject[wanikani.Kanji]) []db.KanjiReading {
	readings := make([]db.KanjiReading, 0)
	for _, reading := range kanji.Data.Readings {
		readings = append(readings, db.KanjiReading{
			Reading: reading.Reading,
			Primary: reading.Primary,
			Type:    reading.Type,
		})
	}

	return readings
}

func createReadingMnemonic(kanji *wanikani.Subject[wanikani.Kanji]) string {
	readingMnemonic := kanji.Data.ReadingMnemonic
	if kanji.Data.ReadingHint != "" {
		readingMnemonic = readingMnemonic + " " + kanji.Data.ReadingHint
	}
	return readingMnemonic
}

func createKanjiMeanings(kanji *wanikani.Subject[wanikani.Kanji]) []db.Meaning {
	meanings := make([]db.Meaning, 0)
	for _, meaning := range kanji.Data.Meanings {
		meanings = append(meanings, db.Meaning{
			Meaning: meaning.Meaning,
			Primary: meaning.Primary,
		})
	}
	for _, auxMeaning := range kanji.Data.AuxiliaryMeanings {
		meanings = append(meanings, db.Meaning{
			Meaning: auxMeaning.Meaning,
			Primary: false,
		})
	}

	return meanings
}
