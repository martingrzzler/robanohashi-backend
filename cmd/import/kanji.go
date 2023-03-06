package main

import (
	"context"
	"fmt"
	"log"
	"robanohashi/cmd/import/wanikani"
	"robanohashi/internal/model"
	"robanohashi/persist"
	"robanohashi/persist/keys"
	"time"

	"github.com/google/uuid"
)

func InsertKanji(ctx context.Context, db *persist.DB, wkKanji *wanikani.Subject[wanikani.Kanji]) {
	id := uuid.New().String()
	meaningMnemonic := model.MeaningMnemonic{
		ID:        id,
		Text:      createKanjiMeaningMnemonic(wkKanji),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		SubjectID: fmt.Sprintf("%d", wkKanji.ID),
	}

	err := db.JSONSet(keys.MeaningMnemonic(id), meaningMnemonic)
	if err != nil {
		log.Fatal(err)
	}

	kanji := model.Kanji{
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
	}

	err = db.JSONSet(keys.Kanji(wkKanji.ID), kanji)
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

func createKanjiReadings(kanji *wanikani.Subject[wanikani.Kanji]) []model.KanjiReading {
	readings := make([]model.KanjiReading, 0)
	for _, reading := range kanji.Data.Readings {
		readings = append(readings, model.KanjiReading{
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

func createKanjiMeanings(kanji *wanikani.Subject[wanikani.Kanji]) []model.Meaning {
	meanings := make([]model.Meaning, 0)
	for _, meaning := range kanji.Data.Meanings {
		meanings = append(meanings, model.Meaning{
			Meaning: meaning.Meaning,
			Primary: meaning.Primary,
		})
	}
	for _, auxMeaning := range kanji.Data.AuxiliaryMeanings {
		meanings = append(meanings, model.Meaning{
			Meaning: auxMeaning.Meaning,
			Primary: false,
		})
	}

	return meanings
}
