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

type MeaningMnemonic struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	SubjectID int    `json:"subject_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
