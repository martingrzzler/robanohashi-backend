package controllers

import (
	"context"
	"net/http"
	"robanohashi/internal/dto"
	"robanohashi/internal/model"
	"robanohashi/persist"
	"robanohashi/persist/keys"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @tags Meaning-Mnemonic
// @summary create a meaning mnemonic for a kanji or vocabulary
// @produce json
// @router /meaning_mnemonic [post]
// @success 201 {object} dto.MeaningMnemonic
// @failure 400 {object} dto.ErrorResponse
// @failure 500 {object} dto.ErrorResponse
// @param request body dto.CreateMeaningMnemonic true "Meaning mnemonic"
// @security firebase
func CreateMeaningMnemonic(c *gin.Context) {
	db := c.MustGet("db").(*persist.DB)
	uid := c.MustGet("uid").(string)

	var body dto.CreateMeaningMnemonic

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	switch body.Object {
	case model.ObjectKanji:
		if _, err := db.GetKanji(context.Background(), body.SubjectID); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Kanji not found"})
			return
		}
	case model.ObjectVocabulary:
		if _, err := db.GetVocabulary(context.Background(), body.SubjectID); err != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Vocabulary not found"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid object"})
		return
	}

	id := uuid.New().String()

	mm := model.MeaningMnemonic{
		ID:        id,
		Text:      body.Text,
		SubjectID: strconv.Itoa(body.SubjectID),
		UserID:    uid,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if err := db.JSONSet(keys.MeaningMnemonic(id), "$", mm); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func VoteMeaningMnemonic(c *gin.Context) {
	db := c.MustGet("db").(*persist.DB)
	uid := c.MustGet("uid").(string)

	var body dto.VoteMeaningMnemonic
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if body.Vote != 1 && body.Vote != -1 {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid vote"})
		return
	}

	if !db.KeyExists(context.Background(), keys.MeaningMnemonic(body.MeaningMnemonicID)) {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not found"})
		return
	}

	switch body.Vote {
	case 1:
		status, err := db.UpvoteMeaningMnemonic(context.Background(), body.MeaningMnemonicID, uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": status})
		return
	case -1:
		status, err := db.DownvoteMeaningMnemonic(context.Background(), body.MeaningMnemonicID, uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": status})
	}
}

func ToggleFavoriteMeaningMnemonic(c *gin.Context) {
	db := c.MustGet("db").(*persist.DB)
	uid := c.MustGet("uid").(string)

	var body dto.ToggleFavoriteMeaningMnemonic
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	if !db.KeyExists(context.Background(), keys.MeaningMnemonic(body.ID)) {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not found"})
		return
	}

	status, err := db.ToggleFavoriteMeaningMnemonic(context.Background(), body.ID, uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}

func UpdateMeaningMnemonic(c *gin.Context) {
	db := c.MustGet("db").(*persist.DB)
	uid := c.MustGet("uid").(string)

	var body dto.UpdateMeaningMnemonic
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	mm, err := db.GetMeaningMnemonic(context.Background(), body.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not found"})
		return
	}

	if mm.UserID != uid {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{Error: "forbidden"})
		return
	}

	if err := db.UpdateMeaningMnemonic(context.Background(), body); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func DeleteMeaningMnemonic(c *gin.Context) {
	db := c.MustGet("db").(*persist.DB)
	uid := c.MustGet("uid").(string)

	var body dto.DeleteMeaningMnemonic

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}

	mm, err := db.GetMeaningMnemonic(context.Background(), body.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "not found"})
		return
	}

	if mm.UserID != uid {
		c.JSON(http.StatusForbidden, dto.ErrorResponse{Error: "forbidden"})
		return
	}

	if err := db.DeleteMeaningMnemonic(context.Background(), body.ID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
