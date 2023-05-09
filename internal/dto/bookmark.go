package dto

import "robanohashi/internal/model"

type ToggleSubjectBookmark struct {
	SubjectID int          `json:"subject_id" binding:"required"`
	Object    model.Object `json:"object" binding:"required"`
}

type BookmarkStatus struct {
	Bookmarked bool `json:"bookmarked"`
}
