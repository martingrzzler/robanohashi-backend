package controllers

import (
	"net/http"
	"robanohashi/internal/dto"
	"robanohashi/persist"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query parameter is required",
		})
		return
	}

	db := c.MustGet("db").(*persist.DB)

	res, err := db.SearchSubjects(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	totalSubject := res.([]interface{})[0].(int64)

	subjects, err := parseSearchResult(res)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"total_count": totalSubject,
		"data":        subjects,
	})
}

func parseSearchResult(result any) ([]dto.SubjectPreview, error) {
	subjects := make([]dto.SubjectPreview, 0)

	for i, subject := range result.([]any)[1:] {
		if i%2 == 0 {
			continue
		}

		preview, err := dto.CreateSubjectPreviewFromRaw(subject.([]any)[1])
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, preview)
	}

	return subjects, nil
}
