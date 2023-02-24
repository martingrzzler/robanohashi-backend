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
	MeaningMnemonicIds  []int               `json:"meaning_mnemonic_ids"`
}

type ContextSentence struct {
	En string `json:"en"`
	Ja string `json:"ja"`
}

type VocabularyReading struct {
	Reading string `json:"reading"`
	Primary bool   `json:"primary"`
	Romaji  string `json:"romaji"`
}
