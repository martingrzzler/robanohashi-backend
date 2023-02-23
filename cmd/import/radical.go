package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"robanohashi/cmd/import/wanikani"
	"robanohashi/keys"
	"robanohashi/model"
)

func InsertRadical(ctx context.Context, cfg Config, wkRadical *wanikani.Subject[wanikani.Radical]) {
	characterImage := make([]byte, 0)
	for _, img := range wkRadical.Data.CharacterImages {
		if img.ContentType == "image/svg+xml" && !img.Metadata.InlineStyles {
			characterImage = FetchCharacterImage(img.URL)
			break
		}
	}

	radical := model.Radical{
		ID:                     wkRadical.ID,
		Object:                 wkRadical.Object,
		Slug:                   wkRadical.Data.Slug,
		Characters:             wkRadical.Data.Characters,
		CharacterImage:         characterImage,
		AmalgamationSubjectIds: wkRadical.Data.AmalgamationSubjectIds,
		Meanings:               createRadicalMeanings(wkRadical),
		MeaningMnemonic:        wkRadical.Data.MeaningMnemonic,
	}

	_, err := cfg.json.JSONSet(keys.Radical(wkRadical.ID), "$", radical)
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

func FetchCharacterImage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
