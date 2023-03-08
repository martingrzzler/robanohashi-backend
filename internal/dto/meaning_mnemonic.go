package dto

import (
	"encoding/json"
	"errors"
	"robanohashi/internal/model"
)

type MeaningMnemonic struct {
	model.MeaningMnemonic
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

type CreateMeaningMnemonic struct {
	Text      string       `json:"text" binding:"required"`
	SubjectID int          `json:"subject_id" binding:"required"`
	Object    model.Object `json:"object" binding:"required"`
}

type VoteMeaningMnemonic struct {
	Vote              int    `json:"vote" binding:"required"`
	MeaningMnemonicID string `json:"meaning_mnemonic_id" binding:"required"`
}

type MeaningMnemonicWithUserInfo struct {
	MeaningMnemonic
	Upvoted   bool `json:"upvoted"`
	Downvoted bool `json:"downvoted"`
}

type UpdateMeaningMnemonic struct {
	ID   string `json:"id" binding:"required"`
	Text string `json:"text" binding:"required"`
}

type DeleteMeaningMnemonic struct {
	ID string `json:"id" binding:"required"`
}

type ToggleFavoriteMeaningMnemonic struct {
	ID string `json:"id" binding:"required"`
}
