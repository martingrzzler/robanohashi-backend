package persist

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"robanohashi/internal/config"
	"robanohashi/internal/model"
	"robanohashi/persist/keys"
)

var db *DB

func TestMain(m *testing.M) {
	cfg := config.New()

	var err error
	db, err = Connect(cfg.RedisURL, cfg.RedisPassword)

	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	err = seed()

	if err != nil {
		log.Fatalf("failed to seed redis: %v", err)
	}

	os.Exit(m.Run())
}

func seed() error {
	db.Client().FlushDB(context.Background())

	err := db.CreateIndices()
	if err != nil {
		return err
	}

	r := model.Radical{
		ID:                     1,
		Object:                 "radical",
		Slug:                   "one",
		Characters:             "一",
		CharacterImage:         "",
		AmalgamationSubjectIds: []int{2},
		Meanings: []model.Meaning{
			{
				Meaning: "one",
				Primary: true,
			},
		},
		MeaningMnemonic: "This is a radical for one",
	}

	k := model.Kanji{
		ID:                     2,
		Object:                 "kanji",
		Slug:                   "one",
		Characters:             "一",
		ReadingMnemonic:        "read it out loud",
		AmalgamationSubjectIds: []int{3},
		Meanings: []model.Meaning{
			{
				Meaning: "one",
				Primary: true,
			},
		},
		Readings: []model.KanjiReading{
			{
				Primary: true,
				Type:    "onyomi",
				Reading: "いち",
			},
		},
		ComponentSubjectIds:       []int{1},
		VisuallySimilarSubjectIds: []int{},
	}

	v := model.Vocabulary{
		ID:                  3,
		Object:              "vocabulary",
		Slug:                "one",
		Characters:          "一",
		ComponentSubjectIds: []int{2},
		Meanings: []model.Meaning{
			{
				Meaning: "one",
				Primary: true,
			},
		},
		ReadingMnemonic: "read it out loud",
		Readings: []model.VocabularyReading{
			{
				Primary: true,
				Romaji:  "ichi",
				Reading: "いち",
			},
		},
		ContextSentences: []model.ContextSentence{
			{
				En:       "I took one step forward",
				Ja:       "私は一歩前に進んだ",
				Hiragana: "わたしはいっぽうまえにすすんだ",
			},
		},
	}

	err = db.JSONSet(keys.Subject(r.ID), "$", r)
	if err != nil {
		return err
	}

	err = db.JSONSet(keys.Subject(k.ID), "$", k)
	if err != nil {
		return err
	}

	err = db.JSONSet(keys.Subject(v.ID), "$", v)
	if err != nil {
		return err
	}

	mm1 := model.MeaningMnemonic{
		ID:          "1",
		Text:        "This is a mnemonic for the meaning of one.",
		VotingCount: 0,
		SubjectID:   "2",
		UserID:      "testuser",
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	mm2 := model.MeaningMnemonic{
		ID:          "2",
		Text:        "This is another mnemonic for the meaning of one.",
		VotingCount: 0,
		SubjectID:   "2",
		UserID:      "testuser",
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	err = db.JSONSet(keys.MeaningMnemonic(mm1.ID), "$", mm1)
	if err != nil {
		return err
	}

	err = db.JSONSet(keys.MeaningMnemonic(mm2.ID), "$", mm2)
	if err != nil {
		return err
	}

	return nil
}
