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
// @success 200 {object} dto.ListResponse[dto.MeaningMnemonicWithUserInfo]
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

	totalCount, mnemonics, err := dto.ParseFTSearchResult[dto.MeaningMnemonic](res)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	_, exists := c.Get("uid")

	if !exists {
		c.JSON(http.StatusOK, dto.ListResponse[dto.MeaningMnemonic]{
			TotalCount: totalCount,
			Data:       mnemonics,
		})
		return
	}

	uid := c.MustGet("uid").(string)

	mnemonicsWithVotes, err := db.ResolveUserVotes(context.Background(), uid, mnemonics)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ListResponse[dto.MeaningMnemonicWithUserInfo]{
		TotalCount: totalCount,
		Data:       mnemonicsWithVotes,
	})
}
