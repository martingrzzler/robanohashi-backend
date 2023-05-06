package dto

import "robanohashi/internal/model"

type Radical struct {
	ID                   int              `json:"id"`
	Object               model.Object     `json:"object"`
	Slug                 string           `json:"slug"`
	Source               string           `json:"source"`
	Characters           string           `json:"characters"`
	CharacterImage       string           `json:"character_image"`
	AmalgamationSubjects []SubjectPreview `json:"amalgamation_subjects"`
	Meanings             []model.Meaning  `json:"meanings"`
	MeaningMnemonic      string           `json:"meaning_mnemonic"`
}
