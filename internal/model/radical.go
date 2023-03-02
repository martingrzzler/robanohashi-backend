package model

type Radical struct {
	ID     int    `json:"id"`
	Object Object `json:"object"`
	Slug   string `json:"slug"`
	// null for some radicals -> use character image (svg)
	Characters             string    `json:"characters"`
	CharacterImage         string    `json:"character_image"`
	AmalgamationSubjectIds []int     `json:"amalgamation_subject_ids"`
	Meanings               []Meaning `json:"meanings"`
	MeaningMnemonic        string    `json:"meaning_mnemonic"`
}

type ResolvedRadical struct {
	ID                   int       `json:"id"`
	Object               Object    `json:"object"`
	Slug                 string    `json:"slug"`
	Characters           string    `json:"characters"`
	CharacterImage       string    `json:"character_image"`
	AmalgamationSubjects []Kanji   `json:"amalgamation_subjects"`
	Meanings             []Meaning `json:"meanings"`
	MeaningMnemonic      string    `json:"meaning_mnemonic"`
}

func (r Radical) GetID() int {
	return r.ID
}

func (r Radical) GetSlug() string {
	return r.Slug
}

func (r Radical) GetObject() Object {
	return r.Object
}

func (r Radical) GetCharacters() string {
	return r.Characters
}

func (r Radical) GetCharacterImage() string {
	return r.CharacterImage
}

func (r Radical) GetMeanings() []Meaning {
	return r.Meanings
}

func (r Radical) GetReadings() []Reading {
	return nil
}
