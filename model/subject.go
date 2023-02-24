package model

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
