package model

type Kanji struct {
	ID                        int            `json:"id"`
	Object                    Object         `json:"object"`
	Characters                string         `json:"characters"`
	Slug                      string         `json:"slug"`
	ReadingMnemonic           string         `json:"reading_mnemonic"`
	AmalgamationSubjectIds    []int          `json:"amalgamation_subject_ids"`
	Meanings                  []Meaning      `json:"meanings"`
	Readings                  []KanjiReading `json:"readings"`
	ComponentSubjectIds       []int          `json:"component_subject_ids"`
	VisuallySimilarSubjectIds []int          `json:"visually_similar_subject_ids"`
}

// contains the actual subject rather than the subject ids
type ResolvedKanji struct {
	ID                      int            `json:"id"`
	Object                  Object         `json:"object"`
	Characters              string         `json:"characters"`
	Slug                    string         `json:"slug"`
	ReadingMnemonic         string         `json:"reading_mnemonic"`
	AmalgamationSubjects    []Vocabulary   `json:"amalgamation_subject_ids"`
	Meanings                []Meaning      `json:"meanings"`
	Readings                []KanjiReading `json:"readings"`
	ComponentSubjects       []Radical      `json:"component_subject_ids"`
	VisuallySimilarSubjects []Kanji        `json:"visually_similar_subject_ids"`
}

func (k Kanji) GetID() int {
	return k.ID
}

func (k Kanji) GetObject() Object {
	return k.Object
}

func (k Kanji) GetSlug() string {
	return k.Slug
}

func (k Kanji) GetCharacters() string {
	return k.Characters
}

func (k Kanji) GetCharacterImage() string {
	return ""
}

func (k Kanji) GetReadings() []Reading {
	readings := make([]Reading, len(k.Readings))
	for i, reading := range k.Readings {
		readings[i] = reading
	}
	return readings
}

func (k Kanji) GetMeanings() []Meaning {
	return k.Meanings
}

type KanjiReading struct {
	Reading string `json:"reading"`
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
}

func (k KanjiReading) GetReading() string {
	return k.Reading
}

func (k KanjiReading) IsPrimary() bool {
	return k.Primary
}
