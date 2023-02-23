package wanikani

import (
	"robanohashi/model"
	"time"
)

type Subject[T Data] struct {
	ID            int          `json:"id"`
	Object        model.Object `json:"object"`
	Url           string       `json:"url"`
	DataUpdatedAt string       `json:"data_updated_at"`
	Data          T            `json:"data"`
}

type Data interface {
	Radical | Kanji | Vocabulary
}

type Radical struct {
	AmalgamationSubjectIds []int `json:"amalgamation_subject_ids"`
	AuxiliaryMeanings      []struct {
		Meaning string `json:"meaning"`
		Type    string `json:"type"`
	} `json:"auxiliary_meanings"`
	Characters      string `json:"characters"`
	CharacterImages []struct {
		URL      string `json:"url"`
		Metadata struct {
			InlineStyles bool `json:"inline_styles"`
		} `json:"metadata"`
		ContentType string `json:"content_type"`
	} `json:"character_images"`
	CreatedAt      time.Time   `json:"created_at"`
	DocumentURL    string      `json:"document_url"`
	HiddenAt       interface{} `json:"hidden_at"`
	LessonPosition int         `json:"lesson_position"`
	Level          int         `json:"level"`
	Meanings       []struct {
		Meaning        string `json:"meaning"`
		Primary        bool   `json:"primary"`
		AcceptedAnswer bool   `json:"accepted_answer"`
	} `json:"meanings"`
	MeaningMnemonic          string `json:"meaning_mnemonic"`
	Slug                     string `json:"slug"`
	SpacedRepetitionSystemID int    `json:"spaced_repetition_system_id"`
}

type Kanji struct {
	AmalgamationSubjectIds []int `json:"amalgamation_subject_ids"`
	AuxiliaryMeanings      []struct {
		Meaning string `json:"meaning"`
		Type    string `json:"type"`
	} `json:"auxiliary_meanings"`
	Characters          string      `json:"characters"`
	ComponentSubjectIds []int       `json:"component_subject_ids"`
	CreatedAt           time.Time   `json:"created_at"`
	DocumentURL         string      `json:"document_url"`
	HiddenAt            interface{} `json:"hidden_at"`
	LessonPosition      int         `json:"lesson_position"`
	Level               int         `json:"level"`
	Meanings            []struct {
		Meaning        string `json:"meaning"`
		Primary        bool   `json:"primary"`
		AcceptedAnswer bool   `json:"accepted_answer"`
	} `json:"meanings"`
	MeaningHint     string `json:"meaning_hint"`
	MeaningMnemonic string `json:"meaning_mnemonic"`
	Readings        []struct {
		Type           string `json:"type"`
		Primary        bool   `json:"primary"`
		AcceptedAnswer bool   `json:"accepted_answer"`
		Reading        string `json:"reading"`
	} `json:"readings"`
	ReadingMnemonic           string `json:"reading_mnemonic"`
	ReadingHint               string `json:"reading_hint"`
	Slug                      string `json:"slug"`
	VisuallySimilarSubjectIds []int  `json:"visually_similar_subject_ids"`
	SpacedRepetitionSystemID  int    `json:"spaced_repetition_system_id"`
}

type Vocabulary struct {
	AuxiliaryMeanings []struct {
		Type    string `json:"type"`
		Meaning string `json:"meaning"`
	} `json:"auxiliary_meanings"`
	Characters          string `json:"characters"`
	ComponentSubjectIds []int  `json:"component_subject_ids"`
	ContextSentences    []struct {
		En string `json:"en"`
		Ja string `json:"ja"`
	} `json:"context_sentences"`
	CreatedAt      time.Time   `json:"created_at"`
	DocumentURL    string      `json:"document_url"`
	HiddenAt       interface{} `json:"hidden_at"`
	LessonPosition int         `json:"lesson_position"`
	Level          int         `json:"level"`
	Meanings       []struct {
		Meaning        string `json:"meaning"`
		Primary        bool   `json:"primary"`
		AcceptedAnswer bool   `json:"accepted_answer"`
	} `json:"meanings"`
	MeaningMnemonic     string   `json:"meaning_mnemonic"`
	PartsOfSpeech       []string `json:"parts_of_speech"`
	PronunciationAudios []struct {
		URL      string `json:"url"`
		Metadata struct {
			Gender           string `json:"gender"`
			SourceID         int    `json:"source_id"`
			Pronunciation    string `json:"pronunciation"`
			VoiceActorID     int    `json:"voice_actor_id"`
			VoiceActorName   string `json:"voice_actor_name"`
			VoiceDescription string `json:"voice_description"`
		} `json:"metadata"`
		ContentType string `json:"content_type"`
	} `json:"pronunciation_audios"`
	Readings []struct {
		Primary        bool   `json:"primary"`
		Reading        string `json:"reading"`
		AcceptedAnswer bool   `json:"accepted_answer"`
	} `json:"readings"`
	ReadingMnemonic          string `json:"reading_mnemonic"`
	Slug                     string `json:"slug"`
	SpacedRepetitionSystemID int    `json:"spaced_repetition_system_id"`
}
