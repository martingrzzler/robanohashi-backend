package controllers

import (
	"context"
	"net/http"
	"robanohashi/internal/dto"
	"robanohashi/persist"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRadical(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	db := c.MustGet("db").(*persist.DB)

	radical, err := db.GetRadical(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "radical not found",
		})
		return
	}

	resolved, err := db.GetRadicalResolved(context.Background(), radical)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to resolve radical",
		})

		return
	}

	c.JSON(http.StatusOK, dto.Radical{
		ID:                   resolved.ID,
		Object:               resolved.Object,
		Slug:                 resolved.Slug,
		Characters:           resolved.Characters,
		CharacterImage:       resolved.CharacterImage,
		AmalgamationSubjects: dto.CreateSubjectPreviews(resolved.AmalgamationSubjects),
		Meanings:             resolved.Meanings,
		MeaningMnemonic:      resolved.MeaningMnemonic,
	})
}
