package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"robanohashi/internal/model"
	"robanohashi/persist/keys"

	"github.com/redis/go-redis/v9"
)

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
