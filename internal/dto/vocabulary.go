package dto

import "robanohashi/internal/model"

type Vocabulary struct {
	ID                int                       `json:"id"`
	Object            model.Object              `json:"object"`
	Slug              string                    `json:"slug"`
	Characters        string                    `json:"characters"`
	ComponentSubjects []SubjectPreview          `json:"component_subjects"`
	Meanings          []model.Meaning           `json:"meanings"`
	ReadingMnemonic   string                    `json:"reading_mnemonic"`
	ContextSentences  []model.ContextSentence   `json:"context_sentences"`
	Readings          []model.VocabularyReading `json:"readings"`
}
