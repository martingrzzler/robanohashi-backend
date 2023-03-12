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

func (db *DB) GetVocabulary(ctx context.Context, id int) (*model.Vocabulary, error) {
	data, err := db.JSONGet(ctx, keys.Subject(id))
	if err != nil {
		return nil, err
	}

	vocab := &model.Vocabulary{}

	err = json.Unmarshal([]byte(data.(string)), vocab)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	if vocab.Object != model.ObjectVocabulary {
		return nil, fmt.Errorf("subject is not a vocabulary")
	}

	return vocab, nil
}

func (db *DB) GetVocabularyResolved(ctx context.Context, vocab *model.Vocabulary) (*dto.Vocabulary, error) {
	pipe := db.rdb.Pipeline()

	componentCmds := make([]*redis.Cmd, len(vocab.ComponentSubjectIds))

	for i, id := range vocab.ComponentSubjectIds {
		componentCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Subject(id))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	resolvedVocab := &dto.Vocabulary{
		ID:                vocab.ID,
		Object:            vocab.Object,
		Slug:              vocab.Slug,
		Characters:        vocab.Characters,
		Meanings:          vocab.Meanings,
		ReadingMnemonic:   vocab.ReadingMnemonic,
		Readings:          vocab.Readings,
		ContextSentences:  vocab.ContextSentences,
		ComponentSubjects: make([]dto.SubjectPreview, len(componentCmds)),
	}

	for i, cmd := range componentCmds {
		kanji := model.Kanji{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &kanji)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedVocab.ComponentSubjects[i] = dto.CreateSubjectPreview(&kanji)
	}

	return resolvedVocab, nil
}
