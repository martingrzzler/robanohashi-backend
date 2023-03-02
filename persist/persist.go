package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"robanohashi/internal/model"
	"robanohashi/persist/keys"
	"strings"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	rdb *redis.Client
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

	err := rdb.Ping(context.Background()).Err()

	if err != nil {
		return nil, err
	}

	return &DB{rdb: rdb}, nil
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
	err = db.createSubjectIndex()
	if err != nil {
		return fmt.Errorf("failed to create subject index: %w", err)
	}

	err = db.createMeaningMnemonicIndex()

	if err != nil {
		return fmt.Errorf("failed to create meaning mnemonic index: %w", err)
	}

	return nil
}

func (db *DB) createMeaningMnemonicIndex() error {
	err := db.rdb.Do(context.Background(),
		"FT.CREATE",
		keys.MeaningMnemonicIndex(),
		"ON", "JSON",
		"PREFIX", "1", "meaning_mnemonic:",
		"SCHEMA",
		"$.subject_id", "AS", "subject_id", "TAG",
	).Err()

	return err
}

func (db *DB) createSubjectIndex() error {
	// Create indices
	err := db.rdb.Do(context.Background(),
		"FT.CREATE",
		keys.SubjectIndex(),
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

func (db *DB) JSONSet(key string, value any) error {
	data, err := json.Marshal(value)

	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	return db.rdb.Do(context.Background(), "JSON.SET", key, "$", string(data)).Err()
}

func (db *DB) Close() {
	db.rdb.Close()
}

func (db *DB) SearchSubjects(search string) (any, error) {

	query := ""
	if len(strings.Split(search, " ")) > 1 {
		query = fmt.Sprintf("@meaning:(%s)", search)
	} else {
		query = fmt.Sprintf("((@characters:{%s}) => { $weight: 2.0 } | (@meaning:(%s)) | (@reading:{%s}) | (@romaji:{%s}))", search, search, search, search)
	}

	return db.rdb.Do(context.Background(), "FT.SEARCH", keys.SubjectIndex(), query).Result()
}

func (db *DB) JSONGet(ctx context.Context, key string) (any, error) {
	data, err := db.rdb.Do(context.Background(), "JSON.GET", key).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to get json: %w", err)
	}

	return data, nil
}

type Result struct {
	Data  any
	Error error
}

func (db *DB) GetKanjiResolved(ctx context.Context, kanji *model.Kanji) (*model.ResolvedKanji, error) {
	pipe := db.rdb.Pipeline()
	componentsCmds := make([]*redis.Cmd, len(kanji.ComponentSubjectIds))
	visuallySimCmds := make([]*redis.Cmd, len(kanji.VisuallySimilarSubjectIds))
	amalgamationCmds := make([]*redis.Cmd, len(kanji.AmalgamationSubjectIds))

	for i, id := range kanji.ComponentSubjectIds {
		componentsCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Radical(id))
	}

	for i, id := range kanji.VisuallySimilarSubjectIds {
		visuallySimCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Kanji(id))
	}

	for i, id := range kanji.AmalgamationSubjectIds {
		amalgamationCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Vocabulary(id))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	resolvedKanji := &model.ResolvedKanji{
		ID:                      kanji.ID,
		Slug:                    kanji.Slug,
		Characters:              kanji.Characters,
		Object:                  kanji.Object,
		Meanings:                kanji.Meanings,
		Readings:                kanji.Readings,
		ReadingMnemonic:         kanji.ReadingMnemonic,
		ComponentSubjects:       make([]model.Radical, len(componentsCmds)),
		VisuallySimilarSubjects: make([]model.Kanji, len(visuallySimCmds)),
		AmalgamationSubjects:    make([]model.Vocabulary, len(amalgamationCmds)),
	}

	for i, cmd := range componentsCmds {
		radical := model.Radical{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &radical)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedKanji.ComponentSubjects[i] = radical
	}

	for i, cmd := range visuallySimCmds {
		kanji := model.Kanji{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &kanji)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedKanji.VisuallySimilarSubjects[i] = kanji
	}

	for i, cmd := range amalgamationCmds {
		vocab := model.Vocabulary{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &vocab)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}
		resolvedKanji.AmalgamationSubjects[i] = vocab
	}

	return resolvedKanji, nil
}

func (db *DB) GetKanji(ctx context.Context, id int) (*model.Kanji, error) {
	data, err := db.JSONGet(ctx, keys.Kanji(id))
	if err != nil {
		return nil, err
	}

	kanji := &model.Kanji{}

	err = json.Unmarshal([]byte(data.(string)), kanji)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return kanji, nil
}

func (db *DB) GetRadical(ctx context.Context, id int) (*model.Radical, error) {
	data, err := db.JSONGet(ctx, keys.Radical(id))
	if err != nil {
		return nil, err
	}

	radical := &model.Radical{}

	err = json.Unmarshal([]byte(data.(string)), radical)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return radical, nil
}

func (db *DB) GetRadicalResolved(ctx context.Context, radical *model.Radical) (*model.ResolvedRadical, error) {
	pipe := db.rdb.Pipeline()

	amalgamationCmds := make([]*redis.Cmd, len(radical.AmalgamationSubjectIds))

	for i, id := range radical.AmalgamationSubjectIds {
		amalgamationCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Kanji(id))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	resolvedRadical := &model.ResolvedRadical{
		ID:                   radical.ID,
		Object:               radical.Object,
		Slug:                 radical.Slug,
		Characters:           radical.Characters,
		CharacterImage:       radical.CharacterImage,
		Meanings:             radical.Meanings,
		MeaningMnemonic:      radical.MeaningMnemonic,
		AmalgamationSubjects: make([]model.Kanji, len(amalgamationCmds)),
	}

	for i, cmd := range amalgamationCmds {
		kanji := model.Kanji{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &kanji)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedRadical.AmalgamationSubjects[i] = kanji
	}

	return resolvedRadical, nil
}
