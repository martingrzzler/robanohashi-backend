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
// @summary get a radical
// @produce json
// @router /radical/{id} [get]
// @success 200 {object} dto.Radical
// @failure 404 {object} dto.ErrorResponse
// @failure 500 {object} dto.ErrorResponse
// @failure 400 {object} dto.ErrorResponse
// @param id path int true "Radical ID"
func GetRadical(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "id must be an integer",
		})
		return
	}

	db := c.MustGet("db").(*persist.DB)

	radical, err := db.GetRadical(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error: "radical not found",
		})
		return
	}

	resolved, err := db.GetRadicalResolved(context.Background(), radical)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "failed to resolve radical",
		})

		return
	}

	c.JSON(http.StatusOK, resolved)
}
