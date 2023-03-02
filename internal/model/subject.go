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
	ID          int    `json:"id"`
	Text        string `json:"text"`
	VotingCount int    `json:"voting_count"`
	// must be string ensure that TAG for the index works as expected
	SubjectID string `json:"subject_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
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
