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

func InsertVocabulary(ctx context.Context, cfg Config, wkVocabulary *wanikani.Subject[wanikani.Vocabulary]) {
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
		Text:      wkVocabulary.Data.MeaningMnemonic,
		CreatedAt: strconv.FormatInt(time.Now().Unix(), 10),
		UpdatedAt: strconv.FormatInt(time.Now().Unix(), 10),
	}

	_, err = cfg.json.JSONSet(keys.MeaningMnemonic(id), "$", meaningMnemonic)
	if err != nil {
		log.Fatal(err)
	}

	vocabulary := db.Vocabulary{
		ID:                  wkVocabulary.ID,
		Object:              wkVocabulary.Object,
		Characters:          wkVocabulary.Data.Characters,
		Slug:                wkVocabulary.Data.Slug,
		ComponentSubjectIds: wkVocabulary.Data.ComponentSubjectIds,
		Meanings:            createVocabMeanings(wkVocabulary),
		ReadingMnemonic:     wkVocabulary.Data.ReadingMnemonic,
		MeaningMnemonicIds:  []int{id},
		Readings:            createVocabReadings(wkVocabulary),
		ContextSentences:    createContextSentences(wkVocabulary),
	}

	_, err = cfg.json.JSONSet(keys.Vocabulary(wkVocabulary.ID), "$", vocabulary)
	if err != nil {
		log.Fatal(err)
	}
}

func createContextSentences(v *wanikani.Subject[wanikani.Vocabulary]) []db.ContextSentence {
	sentences := make([]db.ContextSentence, 0)
	for _, sentence := range v.Data.ContextSentences {
		sentences = append(sentences, db.ContextSentence{
			En: sentence.En,
			Ja: sentence.Ja,
		})
	}

	return sentences
}

func createVocabReadings(v *wanikani.Subject[wanikani.Vocabulary]) []db.VocabularyReading {
	readings := make([]db.VocabularyReading, 0)
	for _, reading := range v.Data.Readings {
		readings = append(readings, db.VocabularyReading{
			Reading: reading.Reading,
			Primary: reading.Primary,
		})
	}

	return readings
}

func createVocabMeanings(v *wanikani.Subject[wanikani.Vocabulary]) []db.Meaning {
	meanings := make([]db.Meaning, 0)
	for _, meaning := range v.Data.Meanings {
		meanings = append(meanings, db.Meaning{
			Meaning: meaning.Meaning,
			Primary: meaning.Primary,
		})
	}
	for _, auxMeaning := range v.Data.AuxiliaryMeanings {
		meanings = append(meanings, db.Meaning{
			Meaning: auxMeaning.Meaning,
			Primary: false,
		})
	}

	return meanings
}
