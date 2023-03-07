package controllers

import (
	"context"
	"net/http"
	"robanohashi/internal/dto"
	"robanohashi/persist"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMeaningMnemonics(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be an integer",
		})
		return
	}

	db := c.MustGet("db").(*persist.DB)

	res, err := db.GetMeaningMnemonicsBySubjectID(context.Background(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	totalCount, mnemonics, err := dto.ParseFTSearchResult[dto.MeaningMnemonic](res)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, exists := c.Get("uid")

	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"total_count": totalCount,
			"data":        mnemonics,
		})
		return
	}

	uid := c.MustGet("uid").(string)

	mnemonicsWithVotes, err := db.ResolveUserVotes(context.Background(), uid, mnemonics)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_count": totalCount,
		"data":        mnemonicsWithVotes,
	})
}
