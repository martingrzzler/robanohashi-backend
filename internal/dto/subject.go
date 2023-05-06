package dto

import "robanohashi/internal/model"

type MnemonicSubject interface {
	GetID() int
	GetObject() model.Object
	GetSlug() string
	GetCharacters() string
	GetReadings() []model.Reading
	GetMeanings() []model.Meaning
	GetComponentSubjects() []SubjectPreview
	GetSource() string
}
