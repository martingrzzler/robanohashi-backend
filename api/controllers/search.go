package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"robanohashi/model"
	"robanohashi/persist"

	"github.com/gin-gonic/gin"
)

type SubjectPreview struct {
	ID             int          `json:"id"`
	Object         model.Object `json:"object"`
	Slug           string       `json:"slug"`
	Characters     string       `json:"characters"`
	CharacterImage string       `json:"character_image"`
	Readings       []string     `json:"readings"`
	Meanings       []string     `json:"meanings"`
}

func Search(c *gin.Context) {
	query := c.Param("query")

	db := c.MustGet("db").(*persist.DB)

	res, err := db.SearchSubjects(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	totalSubject := res.([]interface{})[0].(int64)

	subjects, err := parseSearchResult(res)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"total_count": totalSubject,
		"data":        subjects,
	})
}

func parseSearchResult(result any) ([]SubjectPreview, error) {
	subjects := make([]SubjectPreview, 0)

	for i, subject := range result.([]any)[1:] {
		if i%2 == 0 {
			continue
		}

		preview, err := createSubjectPreviewFromRaw(subject.([]any)[1])
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, preview)
	}

	return subjects, nil
}

// extracts the meaning from the meaning array and places the primary meaning at the beginning
func extractMeanings(ms []model.Meaning) []string {
	meanings := make([]string, 1)
	for _, m := range ms {
		if m.Primary {
			meanings[0] = m.Meaning
			continue
		}
		meanings = append(meanings, m.Meaning)
	}
	return meanings
}

func extractReadings[T model.Reading](rs []T) []string {
	readings := make([]string, 1)
	for _, r := range rs {
		if r.IsPrimary() {
			readings[0] = r.GetReading()
			continue
		}
		readings = append(readings, r.GetReading())
	}
	return readings
}

// @data json string Kanji, Radical or Vocabulary
func createSubjectPreviewFromRaw(data any) (SubjectPreview, error) {
	s, ok := data.(string)
	if !ok {
		return SubjectPreview{}, errors.New("could not convert data to string")
	}

	jsonData := make(map[string]any)

	err := json.Unmarshal([]byte(s), &jsonData)
	if err != nil {
		return SubjectPreview{}, fmt.Errorf("could not unmarshal json: %w", err)
	}

	switch model.Object(jsonData["object"].(string)) {
	case model.ObjectKanji:
		kanji := model.Kanji{}
		err := json.Unmarshal([]byte(s), &kanji)
		if err != nil {
			return SubjectPreview{}, fmt.Errorf("could not unmarshal json: %w", err)
		}
		return SubjectPreview{
			ID:         kanji.ID,
			Object:     kanji.Object,
			Slug:       kanji.Slug,
			Characters: kanji.Characters,
			Meanings:   extractMeanings(kanji.Meanings),
			Readings:   extractReadings(kanji.Readings),
		}, nil

	case model.ObjectRadical:
		radical := model.Radical{}
		err := json.Unmarshal([]byte(s), &radical)
		if err != nil {
			return SubjectPreview{}, fmt.Errorf("could not unmarshal json: %w", err)
		}
		return SubjectPreview{
			ID:             radical.ID,
			Object:         radical.Object,
			Slug:           radical.Slug,
			CharacterImage: radical.CharacterImage,
			Meanings:       extractMeanings(radical.Meanings),
			Characters:     radical.Characters}, nil

	case model.ObjectVocabulary:
		vocabulary := model.Vocabulary{}
		err := json.Unmarshal([]byte(s), &vocabulary)
		if err != nil {
			return SubjectPreview{}, fmt.Errorf("could not unmarshal json: %w", err)
		}
		return SubjectPreview{
			ID:         vocabulary.ID,
			Object:     vocabulary.Object,
			Slug:       vocabulary.Slug,
			Characters: vocabulary.Characters,
			Readings:   extractReadings(vocabulary.Readings),
			Meanings:   extractMeanings(vocabulary.Meanings),
		}, nil
	default:
		panic("detected unsupported object type")

	}
}
