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

func CreateMeaningMnemonic(c *gin.Context) {
	db := c.MustGet("db").(*persist.DB)
	uid := c.MustGet("uid").(string)

	var body dto.CreateMeaningMnemonic

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch body.Object {
	case model.ObjectKanji:
		if _, err := db.GetKanji(context.Background(), body.SubjectID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Kanji not found"})
			return
		}
	case model.ObjectVocabulary:
		if _, err := db.GetVocabulary(context.Background(), body.SubjectID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Vocabulary not found"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid object"})
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

	if err := db.JSONSet(keys.MeaningMnemonic(id), mm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
