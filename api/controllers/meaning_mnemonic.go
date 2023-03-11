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
// @success 201 {object} dto.CreatedResponse
// @failure 400 {object} dto.ErrorResponse
// @failure 500 {object} dto.ErrorResponse
// @param request body dto.CreateMeaningMnemonic true "Meaning mnemonic"
// @security Bearer
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

	c.JSON(http.StatusCreated, dto.CreatedResponse{ID: id})
}

// @tags Meaning-Mnemonic
// @summary vote on a meaning mnemonic
// @produce json
// @router /meaning_mnemonic/vote [post]
// @success 200 {object} dto.StatusResponse
// @failure 400 {object} dto.ErrorResponse
// @failure 500 {object} dto.ErrorResponse
// @failure 404 {object} dto.ErrorResponse
// @param request body dto.VoteMeaningMnemonic true "vote can be 1 or -1"
// @security Bearer
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
		c.JSON(http.StatusOK, dto.StatusResponse{Status: status})
		return
	case -1:
		status, err := db.DownvoteMeaningMnemonic(context.Background(), body.MeaningMnemonicID, uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, dto.StatusResponse{Status: status})
	}
}

// @tags Meaning-Mnemonic
// @summary toggle favorite on a meaning mnemonic
// @produce json
// @router /meaning_mnemonic/toggle_favorite [post]
// @success 200 {object} dto.StatusResponse
// @failure 400 {object} dto.ErrorResponse
// @failure 404 {object} dto.ErrorResponse
// @failure 500 {object} dto.ErrorResponse
// @param request body dto.ToggleFavoriteMeaningMnemonic true "mnemonic id"
// @security Bearer
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

	c.JSON(http.StatusOK, dto.StatusResponse{Status: status})
}

// @tags Meaning-Mnemonic
// @summary update the meaning mnemonic text
// @produce json
// @router /meaning_mnemonic [put]
// @success 200 {object} dto.StatusResponse
// @failure 400 {object} dto.ErrorResponse
// @failure 404 {object} dto.ErrorResponse
// @failure 403 {object} dto.ErrorResponse
// @param request body dto.UpdateMeaningMnemonic true "mnemonic id + text"
// @security Bearer
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

	c.JSON(http.StatusOK, dto.StatusResponse{Status: "ok"})
}

// @tags Meaning-Mnemonic
// @summary delete a meaning mnemonic
// @produce json
// @router /meaning_mnemonic [delete]
// @success 200 {object} dto.StatusResponse
// @failure 400 {object} dto.ErrorResponse
// @failure 404 {object} dto.ErrorResponse
// @failure 403 {object} dto.ErrorResponse
// @failure 500 {object} dto.ErrorResponse
// @security Bearer
// @param request body dto.DeleteMeaningMnemonic true "mnemonic id"
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
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "not found"})
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

	c.JSON(http.StatusOK, dto.StatusResponse{Status: "ok"})
}

// @tags Meaning-Mnemonic
// @summary get all meaning mnemonics marked as favorite
// @produce json
// @router /meaning_mnemonic/favorites [get]
// @success 200 {object} []dto.List[dto.MeaningMnemonicWithUserInfo]
// @failure 500 {object} dto.ErrorResponse
// @security Bearer
func GetFavoriteMeaningMnemonics(c *gin.Context) {
	db := c.MustGet("db").(*persist.DB)
	uid := c.MustGet("uid").(string)

	mms, err := db.GetFavoriteMeaningMnemonics(context.Background(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	resolved, err := db.ResolveUserInfo(context.Background(), uid, mms)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.List[dto.MeaningMnemonicWithUserInfo]{
		TotalCount: int64(len(resolved)),
		Items:      resolved,
	})
}

// @tags Meaning-Mnemonic
// @summary get all meaning mnemonics created by the user
// @produce json
// @router /user/meaning_mnemonics [get]
// @success 200 {object} []dto.List[dto.MeaningMnemonicWithUserInfo]
// @failure 500 {object} dto.ErrorResponse
// @failure 404 {object} dto.ErrorResponse
// @security Bearer
func GetUserMeaningMnemonics(c *gin.Context) {
	db := c.MustGet("db").(*persist.DB)
	uid := c.MustGet("uid").(string)

	res, err := db.GetMeaningMnemonicsByUser(context.Background(), uid)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Error: err.Error()})
		return
	}

	resolved, err := db.ResolveUserInfo(context.Background(), uid, res.Items)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.List[dto.MeaningMnemonicWithUserInfo]{
		TotalCount: res.TotalCount,
		Items:      resolved,
	})
}
