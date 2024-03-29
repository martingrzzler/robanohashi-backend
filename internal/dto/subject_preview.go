package dto

import (
	"encoding/json"
	"errors"
	"fmt"
	"robanohashi/internal/model"
)

type SubjectPreview struct {
	ID             int          `json:"id"`
	Object         model.Object `json:"object"`
	Slug           string       `json:"slug"`
	Source         string       `json:"source"`
	Characters     string       `json:"characters"`
	CharacterImage string       `json:"character_image"`
	Readings       []string     `json:"readings"`
	Meanings       []string     `json:"meanings"`
}

func CreateSubjectPreview(subject model.Subject) SubjectPreview {
	return SubjectPreview{
		ID:             subject.GetID(),
		Slug:           subject.GetSlug(),
		Object:         subject.GetObject(),
		Source:         subject.GetSource(),
		Characters:     subject.GetCharacters(),
		CharacterImage: subject.GetCharacterImage(),
		Readings:       extractReadings(subject.GetReadings()),
		Meanings:       extractMeanings(subject.GetMeanings()),
	}
}

func CreateSubjectPreviews[T model.Subject](subjects []T) []SubjectPreview {
	res := make([]SubjectPreview, 0)

	for _, subject := range subjects {
		res = append(res, CreateSubjectPreview(subject))
	}

	return res
}

// extracts the meaning from the meaning array and places the primary meaning at the beginning
func extractMeanings(ms []model.Meaning) []string {
	meanings := make([]string, 0)
	for _, m := range ms {
		if m.Primary {
			meanings = append([]string{m.Meaning}, meanings...)
			continue
		}
		meanings = append(meanings, m.Meaning)
	}
	return meanings
}

func extractReadings[T model.Reading](rs []T) []string {
	readings := make([]string, 0)
	for _, r := range rs {
		if r.IsPrimary() {
			readings = append([]string{r.GetReading()}, readings...)
			continue
		}
		readings = append(readings, r.GetReading())
	}
	return readings
}

func (p SubjectPreview) UnmarshalRaw(data any) (SubjectPreview, error) {
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
		return CreateSubjectPreview(kanji), nil

	case model.ObjectRadical:
		radical := model.Radical{}
		err := json.Unmarshal([]byte(s), &radical)
		if err != nil {
			return SubjectPreview{}, fmt.Errorf("could not unmarshal json: %w", err)
		}
		return CreateSubjectPreview(radical), nil

	case model.ObjectVocabulary:
		vocabulary := model.Vocabulary{}
		err := json.Unmarshal([]byte(s), &vocabulary)
		if err != nil {
			return SubjectPreview{}, fmt.Errorf("could not unmarshal json: %w", err)
		}
		return CreateSubjectPreview(vocabulary), nil
	default:
		panic("detected unsupported object type")

	}
}
