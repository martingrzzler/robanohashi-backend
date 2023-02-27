package controllers

import (
	"robanohashi/model"

	"github.com/gin-gonic/gin"
)

type Kanji struct {
	ID                        int                     `json:"id"`
	Object                    model.Object            `json:"object"`
	Characters                string                  `json:"characters"`
	Slug                      string                  `json:"slug"`
	ReadingMnemonic           string                  `json:"reading_mnemonic"`
	AmalgamationSubjects      []SubjectPreview        `json:"amalgamation_subjects"`
	Meanings                  []model.Meaning         `json:"meanings"`
	Readings                  []model.KanjiReading    `json:"readings"`
	ComponentSubjects         []SubjectPreview        `json:"component_subjects"`
	VisuallySimilarSubjectIds []SubjectPreview        `json:"visually_similar_subject"`
	MeaningMnemonicIds        []model.MeaningMnemonic `json:"meaning_mnemonic_ids"`
}

func GetKanji(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "id must be an integer",
	// 	})
	// 	return
	// }

	// db := c.MustGet("db").(*persist.DB)

	// res, err := db.GetKanji(id)
}
