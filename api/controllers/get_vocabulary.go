package controllers

import (
	"context"
	"net/http"
	"robanohashi/internal/dto"
	"robanohashi/persist"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @tags Subject
// @summary get a vocabulary
// @produce json
// @router /vocabulary/{id} [get]
// @success 200 {object} dto.Vocabulary
// @failure 404 {object} dto.ErrorResponse
// @failure 500 {object} dto.ErrorResponse
// @failure 400 {object} dto.ErrorResponse
// @param id path int true "Vocabulary ID"
func GetVocabulary(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "id must be an integer",
		})
		return
	}

	db := c.MustGet("db").(*persist.DB)

	vocabulary, err := db.GetVocabulary(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "vocabulary not found",
		})
		return
	}

	resolved, err := db.GetVocabularyResolved(context.Background(), vocabulary)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "failed to resolve vocabulary",
		})

		return
	}

	c.JSON(http.StatusOK, dto.Vocabulary{
		ID:                resolved.ID,
		Object:            resolved.Object,
		Slug:              resolved.Slug,
		Characters:        resolved.Characters,
		ComponentSubjects: dto.CreateSubjectPreviews(resolved.ComponentSubjects),
		Meanings:          resolved.Meanings,
		ReadingMnemonic:   resolved.ReadingMnemonic,
		ContextSentences:  resolved.ContextSentences,
		Readings:          resolved.Readings,
	})
}
