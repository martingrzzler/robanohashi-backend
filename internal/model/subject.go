package model

import (
	"encoding/json"
	"errors"
)

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

type NonHumanUserID string

const (
	NonHumanUserIDWaniKani    NonHumanUserID = "wanikani"
	NonHumanUserIDAIGenerated NonHumanUserID = "ai_generated"
)

type MeaningMnemonic struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	VotingCount int    `json:"voting_count"`
	// must be string ensure that TAG for the index works as expected
	SubjectID string `json:"subject_id"`
	UserID    string `json:"user_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (m MeaningMnemonic) UnmarshalRaw(data any) (MeaningMnemonic, error) {
	s, ok := data.(string)
	if !ok {
		return MeaningMnemonic{}, errors.New("could not convert data to string")
	}

	mm := MeaningMnemonic{}

	err := json.Unmarshal([]byte(s), &mm)
	if err != nil {
		return MeaningMnemonic{}, err
	}
	return mm, nil
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
