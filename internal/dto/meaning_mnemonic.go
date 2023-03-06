package dto

import "robanohashi/internal/model"

type CreateMeaningMnemonic struct {
	Text      string       `json:"text" binding:"required"`
	SubjectID int          `json:"subject_id" binding:"required"`
	Object    model.Object `json:"object" binding:"required"`
}

type VoteMeaningMnemonic struct {
	Vote              int    `json:"vote" binding:"required"`
	MeaningMnemonicID string `json:"meaning_mnemonic_id" binding:"required"`
}
