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
// @summary get a kanji
// @produce json
// @router /kanji/{id} [get]
// @success 200 {object} dto.Kanji
// @failure 404 {object} dto.ErrorResponse
// @failure 500 {object} dto.ErrorResponse
// @failure 400 {object} dto.ErrorResponse
// @param id path int true "Kanji ID"
func GetKanji(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "id must be an integer",
		})
		return
	}

	db := c.MustGet("db").(*persist.DB)

	kanji, err := db.GetKanji(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "kanji not found",
		})
		return
	}

	resolved, err := db.GetKanjiResolved(context.Background(), kanji)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "failed to resolve kanji",
		})
		return
	}

	c.JSON(http.StatusOK, resolved)
}
