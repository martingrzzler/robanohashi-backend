package persist

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"robanohashi/model"
	"robanohashi/persist/keys"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
)

type DB struct {
	rdb         *redis.Client
	jsonHandler *rejson.Handler
}

func (db *DB) JSONHandler() *rejson.Handler {
	return db.jsonHandler
}

func (db *DB) Client() *redis.Client {
	return db.rdb
}

func Connect() (*DB, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		return nil, err
	}

	handler := rejson.NewReJSONHandler()
	handler.SetGoRedisClient(rdb)

	return &DB{rdb: rdb, jsonHandler: handler}, nil
}

func (db *DB) CreateIndices() error {
	res, err := db.rdb.Do(context.Background(), "FT._LIST").Result()
	if err != nil {
		return err
	}

	for _, index := range res.([]interface{}) {
		err := db.rdb.Do(context.Background(), "FT.DROPINDEX", index).Err()
		if err != nil {
			return err
		}
	}

	// Create indices
	err = db.rdb.Do(context.Background(),
		"FT.CREATE",
		keys.SearchIndex(),
		"ON", "JSON",
		"PREFIX", "3", "kanji:", "radical:", "vocabulary:",
		"SCHEMA",
		"$.characters", "AS", "characters", "TAG",
		"$.meanings.*.meaning", "AS", "meaning", "TEXT",
		"$.readings.*.reading", "AS", "reading", "TAG",
		"$.readings.*.romaji", "AS", "romaji", "TAG",
	).Err()

	return err

}

func (db *DB) Close() {
	db.rdb.Close()
}

type SubjectPreview struct {
	ID             int          `json:"id"`
	Object         model.Object `json:"object"`
	Slug           string       `json:"slug"`
	Characters     string       `json:"characters"`
	CharacterImage string       `json:"character_image"`
	Readings       []string     `json:"readings"`
	Meanings       []string     `json:"meanings"`
}

func (db *DB) SearchSubjects(search string) ([]SubjectPreview, error) {

	query := ""
	if len(strings.Split(search, " ")) > 1 {
		query = fmt.Sprintf("@meaning:(%s)", search)
	} else {
		query = fmt.Sprintf("((@characters:{%s}) => { $weight: 2.0 } | (@meaning:(%s)) | (@reading:{%s}) | (@romaji:{%s}))", search, search, search, search)
	}

	res, err := db.rdb.Do(context.Background(), "FT.SEARCH", keys.SearchIndex(), query).Result()
	if err != nil {
		return nil, err
	}

	// totalSubjects := res.([]any)[0].(int64)

	subjects := make([]SubjectPreview, 0)

	for i, subject := range res.([]any)[1:] {
		if i%2 == 0 {
			continue
		}

		preview, err := NewSubjectPreviewFromSearch(subject.([]any)[1])
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

func NewSubjectPreviewFromSearch(data any) (SubjectPreview, error) {
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
