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
// @summary get meaning mnemonics optionally with user data if authenticated
// @produce json
// @router /subject/{id}/meaning_mnemonics [get]
// @success 200 {object} dto.List[dto.MeaningMnemonicWithUserInfo]
// @failure 500 {object} dto.ErrorResponse
// @failure 400 {object} dto.ErrorResponse
// @param id path int true "Subject ID vocabulary or kanji"
func GetMeaningMnemonics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "id must be an integer",
		})
		return
	}

	db := c.MustGet("db").(*persist.DB)

	res, err := db.GetMeaningMnemonicsBySubjectID(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Something went wrong",
		})
		return
	}

	_, exists := c.Get("uid")

	if !exists {
		c.JSON(http.StatusOK, res)
		return
	}

	uid := c.MustGet("uid").(string)

	mnemonicsWithUserInfo, err := db.ResolveUserInfo(context.Background(), uid, res.Items)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.List[dto.MeaningMnemonicWithUserInfo]{
		TotalCount: res.TotalCount,
		Items:      mnemonicsWithUserInfo,
	})
}
