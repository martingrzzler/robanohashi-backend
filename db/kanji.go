package db

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
	MeaningMnemonicIds        []int          `json:"meaning_mnemonic_ids"`
}

type MeaningMnemonic struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type KanjiReading struct {
	Reading string `json:"reading"`
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
}
