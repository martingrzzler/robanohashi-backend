package dto

import "robanohashi/internal/model"

type Kanji struct {
	ID                      int                  `json:"id"`
	Object                  model.Object         `json:"object"`
	Characters              string               `json:"characters"`
	Slug                    string               `json:"slug"`
	ReadingMnemonic         string               `json:"reading_mnemonic"`
	AmalgamationSubjects    []SubjectPreview     `json:"amalgamation_subjects"`
	Meanings                []model.Meaning      `json:"meanings"`
	Readings                []model.KanjiReading `json:"readings"`
	ComponentSubjects       []SubjectPreview     `json:"component_subjects"`
	VisuallySimilarSubjects []SubjectPreview     `json:"visually_similar_subjects"`
}

func (k Kanji) GetID() int {
	return k.ID
}

func (k Kanji) GetObject() model.Object {
	return k.Object
}

func (k Kanji) GetCharacters() string {
	return k.Characters
}

func (k Kanji) GetSlug() string {
	return k.Slug
}

func (k Kanji) GetReadings() []model.Reading {
	readings := make([]model.Reading, len(k.Readings))
	for i, reading := range k.Readings {
		readings[i] = reading
	}
	return readings
}

func (k Kanji) GetMeanings() []model.Meaning {
	return k.Meanings
}

func (k Kanji) GetComponentSubjects() []SubjectPreview {
	return k.ComponentSubjects
}
