package model

type Radical struct {
	ID     int    `json:"id"`
	Object Object `json:"object"`
	Slug   string `json:"slug"`
	// null for some radicals -> use character image (svg)
	Characters             string    `json:"characters"`
	CharacterImage         []byte    `json:"character_image"`
	AmalgamationSubjectIds []int     `json:"amalgamation_subject_ids"`
	Meanings               []Meaning `json:"meanings"`
	MeaningMnemonic        string    `json:"meaning_mnemonic"`
}
