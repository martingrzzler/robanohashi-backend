package main

import (
	"context"
	"log"
	"regexp"
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
		CharacterImage:         setStrokeToWhite(wkRadical.Data.CharacterSvgImage),
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

func setStrokeToWhite(svg string) string {
	re := regexp.MustCompile("stroke:#000;")

	return re.ReplaceAllString(svg, "stroke:#fff;")
}
