package controllers

import (
	"context"
	"net/http"
	"robanohashi/internal/dto"
	"robanohashi/persist"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetVocabulary(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	db := c.MustGet("db").(*persist.DB)

	vocabulary, err := db.GetVocabulary(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "vocabulary not found",
		})
		return
	}

	resolved, err := db.GetVocabularyResolved(context.Background(), vocabulary)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to resolve vocabulary",
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
