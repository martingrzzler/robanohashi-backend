package model

type Vocabulary struct {
	ID                  int                 `json:"id"`
	Object              Object              `json:"object"`
	Slug                string              `json:"slug"`
	Characters          string              `json:"characters"`
	ComponentSubjectIds []int               `json:"component_subject_ids"`
	Meanings            []Meaning           `json:"meanings"`
	ReadingMnemonic     string              `json:"reading_mnemonic"`
	ContextSentences    []ContextSentence   `json:"context_sentences"`
	Readings            []VocabularyReading `json:"readings"`
}

type ResolvedVocabulary struct {
	ID                int                 `json:"id"`
	Object            Object              `json:"object"`
	Slug              string              `json:"slug"`
	Characters        string              `json:"characters"`
	ComponentSubjects []Kanji             `json:"component_subjects"`
	Meanings          []Meaning           `json:"meanings"`
	ReadingMnemonic   string              `json:"reading_mnemonic"`
	ContextSentences  []ContextSentence   `json:"context_sentences"`
	Readings          []VocabularyReading `json:"readings"`
}

func (v Vocabulary) GetID() int {
	return v.ID
}

func (v Vocabulary) GetSlug() string {
	return v.Slug
}

func (v Vocabulary) GetObject() Object {
	return v.Object
}

func (v Vocabulary) GetCharacters() string {
	return v.Characters
}

func (v Vocabulary) GetCharacterImage() string {
	return ""
}

func (v Vocabulary) GetMeanings() []Meaning {
	return v.Meanings
}

func (v Vocabulary) GetReadings() []Reading {
	readings := make([]Reading, len(v.Readings))

	for i, r := range v.Readings {
		readings[i] = r
	}

	return readings
}

type ContextSentence struct {
	En       string `json:"en"`
	Ja       string `json:"ja"`
	Hiragana string `json:"hiragana"`
}

type VocabularyReading struct {
	Reading string `json:"reading"`
	Primary bool   `json:"primary"`
	Romaji  string `json:"romaji"`
}

func (v VocabularyReading) GetReading() string {
	return v.Reading
}

func (v VocabularyReading) IsPrimary() bool {
	return v.Primary
}
