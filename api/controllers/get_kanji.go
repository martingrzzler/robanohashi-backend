package controllers

import (
	"context"
	"net/http"
	"robanohashi/internal/dto"
	"robanohashi/persist"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetKanji(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	db := c.MustGet("db").(*persist.DB)

	kanji, err := db.GetKanji(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "kanji not found",
		})
		return
	}

	resolved, err := db.GetKanjiResolved(context.Background(), kanji)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to resolve kanji",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Kanji{
		ID:                      resolved.ID,
		Object:                  resolved.Object,
		Slug:                    resolved.Slug,
		Characters:              resolved.Characters,
		Meanings:                resolved.Meanings,
		Readings:                resolved.Readings,
		ReadingMnemonic:         resolved.ReadingMnemonic,
		VisuallySimilarSubjects: dto.CreateSubjectPreviews(resolved.VisuallySimilarSubjects),
		ComponentSubjects:       dto.CreateSubjectPreviews(resolved.ComponentSubjects),
		AmalgamationSubjects:    dto.CreateSubjectPreviews(resolved.AmalgamationSubjects),
	})
}
