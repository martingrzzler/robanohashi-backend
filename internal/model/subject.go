package model

type Reading interface {
	GetReading() string
	IsPrimary() bool
}

type Object string

const (
	ObjectKanji      Object = "kanji"
	ObjectRadical    Object = "radical"
	ObjectVocabulary Object = "vocabulary"
)

type Meaning struct {
	Meaning string `json:"meaning"`
	Primary bool   `json:"primary"`
}

type MeaningMnemonic struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	// must be string ensure that TAG for the index works as expected
	SubjectID string `json:"subject_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Subject interface {
	GetID() int
	GetObject() Object
	GetSlug() string
	GetCharacters() string
	GetCharacterImage() string
	GetReadings() []Reading
	GetMeanings() []Meaning
}
