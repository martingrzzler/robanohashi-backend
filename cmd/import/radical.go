package main

import (
	"context"
	"log"
	"robanohashi/cmd/import/wanikani"
	"robanohashi/model"
	"robanohashi/persist"
	"robanohashi/persist/keys"
)

func InsertRadical(ctx context.Context, db *persist.DB, wkRadical *wanikani.Subject[wanikani.Radical]) {
	radical := model.Radical{
		ID:                     wkRadical.ID,
		Object:                 wkRadical.Object,
		Slug:                   wkRadical.Data.Slug,
		Characters:             wkRadical.Data.Characters,
		CharacterImage:         wkRadical.Data.CharacterSvgImage,
		AmalgamationSubjectIds: wkRadical.Data.AmalgamationSubjectIds,
		Meanings:               createRadicalMeanings(wkRadical),
		MeaningMnemonic:        wkRadical.Data.MeaningMnemonic,
	}

	_, err := db.JSONHandler().JSONSet(keys.Radical(wkRadical.ID), "$", radical)
	if err != nil {
		log.Fatal(err)
	}
}

func createRadicalMeanings(wkRadical *wanikani.Subject[wanikani.Radical]) []model.Meaning {
	meanings := make([]model.Meaning, len(wkRadical.Data.Meanings))
	for i, meaning := range wkRadical.Data.Meanings {
		meanings[i] = model.Meaning{
			Meaning: meaning.Meaning,
			Primary: meaning.Primary,
		}
	}
	return meanings
}
