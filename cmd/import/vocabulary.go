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

	"github.com/gojp/kana"
	"github.com/google/uuid"
)

func InsertVocabulary(ctx context.Context, db *persist.DB, wkVocabulary *wanikani.Subject[wanikani.Vocabulary]) {
	id := uuid.New().String()

	meaningMnemonic := model.MeaningMnemonic{
		ID:        id,
		UserID:    string(model.NonHumanUserIDWaniKani),
		Text:      wkVocabulary.Data.MeaningMnemonic,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		SubjectID: fmt.Sprintf("%d", wkVocabulary.ID),
	}

	err := db.JSONSet(keys.MeaningMnemonic(id), "$", meaningMnemonic)
	if err != nil {
		log.Fatal(err)
	}

	vocabulary := model.Vocabulary{
		ID:                  wkVocabulary.ID,
		Object:              wkVocabulary.Object,
		Characters:          wkVocabulary.Data.Characters,
		Slug:                wkVocabulary.Data.Slug,
		ComponentSubjectIds: wkVocabulary.Data.ComponentSubjectIds,
		Meanings:            createVocabMeanings(wkVocabulary),
		ReadingMnemonic:     wkVocabulary.Data.ReadingMnemonic,
		Readings:            createVocabReadings(wkVocabulary),
		ContextSentences:    createContextSentences(wkVocabulary),
	}

	err = db.JSONSet(keys.Subject(wkVocabulary.ID), "$", vocabulary)
	if err != nil {
		log.Fatal(err)
	}
}

func createContextSentences(v *wanikani.Subject[wanikani.Vocabulary]) []model.ContextSentence {
	sentences := make([]model.ContextSentence, 0)
	for _, sentence := range v.Data.ContextSentences {
		sentences = append(sentences, model.ContextSentence{
			En:       sentence.En,
			Ja:       sentence.Ja,
			Hiragana: sentence.Hiragana,
		})
	}

	return sentences
}

func createVocabReadings(v *wanikani.Subject[wanikani.Vocabulary]) []model.VocabularyReading {
	readings := make([]model.VocabularyReading, 0)
	for _, reading := range v.Data.Readings {
		readings = append(readings, model.VocabularyReading{
			Reading: reading.Reading,
			Primary: reading.Primary,
			Romaji:  kana.KanaToRomaji(reading.Reading),
		})
	}

	return readings
}

func createVocabMeanings(v *wanikani.Subject[wanikani.Vocabulary]) []model.Meaning {
	meanings := make([]model.Meaning, 0)
	for _, meaning := range v.Data.Meanings {
		meanings = append(meanings, model.Meaning{
			Meaning: meaning.Meaning,
			Primary: meaning.Primary,
		})
	}
	for _, auxMeaning := range v.Data.AuxiliaryMeanings {
		meanings = append(meanings, model.Meaning{
			Meaning: auxMeaning.Meaning,
			Primary: false,
		})
	}

	return meanings
}
