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
