package dto

import (
	"robanohashi/internal/model"
)

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
	ID          string          `json:"id"`
	Text        string          `json:"text"`
	VotingCount int             `json:"voting_count"`
	Subject     MnemonicSubject `json:"subject"`
	UserID      string          `json:"user_id"`
	CreatedAt   int64           `json:"created_at"`
	UpdatedAt   int64           `json:"updated_at"`
	Upvoted     bool            `json:"upvoted"`
	Downvoted   bool            `json:"downvoted"`
	Favorite    bool            `json:"favorite"`
	Me          bool            `json:"me"`
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
