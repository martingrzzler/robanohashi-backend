package controllers

import (
	"context"
	"net/http"
	"robanohashi/internal/dto"
	"robanohashi/internal/model"
	"robanohashi/persist"
	"robanohashi/persist/keys"

	"github.com/gin-gonic/gin"
)

// @tags Subject
// @summary toggle bookmark for a subject
// @produce json
// @router /subject/bookmark [post]
// @success 200 {object} string
// @failure 500 {object} dto.ErrorResponse
// @failure 400 {object} dto.ErrorResponse
// @failure 404 {object} dto.ErrorResponse
// @param request body dto.ToggleSubjectBookmark true "toggle subject bookmark"
// @security Bearer
func ToggleSubjectBookmarked(c *gin.Context) {
	db := c.MustGet("db").(*persist.DB)
	uid := c.MustGet("uid").(string)

	var body dto.ToggleSubjectBookmark

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	subjectKey := ""

	switch body.Object {
	case model.ObjectKanji:
		if !db.KeyExists(context.Background(), keys.Kanji(body.SubjectID)) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Kanji not found"})
			return
		}

		subjectKey = keys.Kanji(body.SubjectID)

	case model.ObjectVocabulary:
		if !db.KeyExists(context.Background(), keys.Vocabulary(body.SubjectID)) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Vocabulary not found"})
			return
		}

		subjectKey = keys.Vocabulary(body.SubjectID)

	case model.ObjectRadical:
		if !db.KeyExists(context.Background(), keys.Radical(body.SubjectID)) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "Radical not found"})
			return
		}

		subjectKey = keys.Radical(body.SubjectID)
	default:
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid object"})
		return
	}

	status, err := db.ToggleSubjectBookmarked(context.Background(), subjectKey, uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.StatusResponse{Status: status})
	return
}

// @tags Subject
// @summary get user bookmarked subjects
// @produce json
// @router /user/bookmarks [get]
// @success 200 {object} []dto.List[dto.SubjectPreview]
// @failure 500 {object} dto.ErrorResponse
// @security Bearer
func GetUserBookmarkedSubjects(c *gin.Context) {
	db := c.MustGet("db").(*persist.DB)
	uid := c.MustGet("uid").(string)

	subjects, err := db.GetUserBookmarkedSubjects(context.Background(), uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, subjects)
	return
}
