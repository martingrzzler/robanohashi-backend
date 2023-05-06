package dto

import "robanohashi/internal/model"

type Vocabulary struct {
	ID                int                       `json:"id"`
	Object            model.Object              `json:"object"`
	Slug              string                    `json:"slug"`
	Source            string                    `json:"source"`
	Characters        string                    `json:"characters"`
	ComponentSubjects []SubjectPreview          `json:"component_subjects"`
	Meanings          []model.Meaning           `json:"meanings"`
	ReadingMnemonic   string                    `json:"reading_mnemonic"`
	ContextSentences  []model.ContextSentence   `json:"context_sentences"`
	Readings          []model.VocabularyReading `json:"readings"`
}

func (v Vocabulary) GetSource() string {
	return v.Source
}

func (v Vocabulary) GetID() int {
	return v.ID
}

func (v Vocabulary) GetSlug() string {
	return v.Slug
}

func (v Vocabulary) GetObject() model.Object {
	return v.Object
}

func (v Vocabulary) GetCharacters() string {
	return v.Characters
}

func (v Vocabulary) GetMeanings() []model.Meaning {
	return v.Meanings
}

func (v Vocabulary) GetReadings() []model.Reading {
	readings := make([]model.Reading, len(v.Readings))

	for i, r := range v.Readings {
		readings[i] = r
	}

	return readings
}

func (v Vocabulary) GetComponentSubjects() []SubjectPreview {
	return v.ComponentSubjects
}
