package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"robanohashi/internal/dto"
	"robanohashi/internal/model"
	"robanohashi/persist/keys"

	"github.com/redis/go-redis/v9"
)

func (db *DB) GetKanji(ctx context.Context, id int) (*model.Kanji, error) {
	data, err := db.JSONGet(ctx, keys.Subject(id))
	if err != nil {
		return nil, err
	}

	kanji := &model.Kanji{}

	err = json.Unmarshal([]byte(data.(string)), kanji)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	if kanji.Object != model.ObjectKanji {
		return nil, fmt.Errorf("subject is not a kanji")
	}

	return kanji, nil
}

func (db *DB) GetKanjiResolved(ctx context.Context, kanji *model.Kanji) (*dto.Kanji, error) {
	pipe := db.rdb.Pipeline()
	componentsCmds := make([]*redis.Cmd, len(kanji.ComponentSubjectIds))
	visuallySimCmds := make([]*redis.Cmd, len(kanji.VisuallySimilarSubjectIds))
	amalgamationCmds := make([]*redis.Cmd, len(kanji.AmalgamationSubjectIds))

	for i, id := range kanji.ComponentSubjectIds {
		componentsCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Subject(id))
	}

	for i, id := range kanji.VisuallySimilarSubjectIds {
		visuallySimCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Subject(id))
	}

	for i, id := range kanji.AmalgamationSubjectIds {
		amalgamationCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Subject(id))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	resolvedKanji := &dto.Kanji{
		ID:                      kanji.ID,
		Slug:                    kanji.Slug,
		Characters:              kanji.Characters,
		Object:                  kanji.Object,
		Meanings:                kanji.Meanings,
		Readings:                kanji.Readings,
		ReadingMnemonic:         kanji.ReadingMnemonic,
		ComponentSubjects:       make([]dto.SubjectPreview, len(componentsCmds)),
		VisuallySimilarSubjects: make([]dto.SubjectPreview, len(visuallySimCmds)),
		AmalgamationSubjects:    make([]dto.SubjectPreview, len(amalgamationCmds)),
	}

	for i, cmd := range componentsCmds {
		radical := model.Radical{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &radical)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedKanji.ComponentSubjects[i] = dto.CreateSubjectPreview(radical)
	}

	for i, cmd := range visuallySimCmds {
		kanji := model.Kanji{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &kanji)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedKanji.VisuallySimilarSubjects[i] = dto.CreateSubjectPreview(kanji)
	}

	for i, cmd := range amalgamationCmds {
		vocab := model.Vocabulary{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &vocab)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}
		resolvedKanji.AmalgamationSubjects[i] = dto.CreateSubjectPreview(vocab)
	}

	return resolvedKanji, nil
}
