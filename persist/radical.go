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

func (db *DB) GetRadicalResolved(ctx context.Context, radical *model.Radical) (*dto.Radical, error) {
	pipe := db.rdb.Pipeline()

	amalgamationCmds := make([]*redis.Cmd, len(radical.AmalgamationSubjectIds))

	for i, id := range radical.AmalgamationSubjectIds {
		amalgamationCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Kanji(id))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	resolvedRadical := &dto.Radical{
		ID:                   radical.ID,
		Object:               radical.Object,
		Slug:                 radical.Slug,
		Source:               radical.Source.String(),
		Characters:           radical.Characters,
		CharacterImage:       radical.CharacterImage,
		Meanings:             radical.Meanings,
		MeaningMnemonic:      radical.MeaningMnemonic,
		AmalgamationSubjects: make([]dto.SubjectPreview, len(amalgamationCmds)),
	}

	for i, cmd := range amalgamationCmds {
		kanji := model.Kanji{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &kanji)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedRadical.AmalgamationSubjects[i] = dto.CreateSubjectPreview(kanji)
	}

	return resolvedRadical, nil
}
