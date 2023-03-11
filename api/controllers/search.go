package controllers

import (
	"context"
	"net/http"
	"robanohashi/internal/dto"
	"robanohashi/persist"

	"github.com/gin-gonic/gin"
)

// @tags Search
// @summary search for subjects
// @produce json
// @router /search [get]
// @success 200 {object} dto.ListResponse[dto.SubjectPreview]
// @failure 500 {object} dto.ErrorResponse
// @failure 400 {object} dto.ErrorResponse
// @param query query string true "Search query"
func Search(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "query parameter is required",
		})
		return
	}

	db := c.MustGet("db").(*persist.DB)

	data, err := db.SearchSubjects(context.Background(), query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
