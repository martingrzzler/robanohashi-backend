package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"robanohashi/internal/model"
	"robanohashi/persist/keys"

	"github.com/redis/go-redis/v9"
)

func (db *DB) GetVocabulary(ctx context.Context, id int) (*model.Vocabulary, error) {
	data, err := db.JSONGet(ctx, keys.Vocabulary(id))
	if err != nil {
		return nil, err
	}

	vocab := &model.Vocabulary{}

	err = json.Unmarshal([]byte(data.(string)), vocab)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return vocab, nil
}

func (db *DB) GetVocabularyResolved(ctx context.Context, vocab *model.Vocabulary) (*model.ResolvedVocabulary, error) {
	pipe := db.rdb.Pipeline()

	componentCmds := make([]*redis.Cmd, len(vocab.ComponentSubjectIds))

	for i, id := range vocab.ComponentSubjectIds {
		componentCmds[i] = pipe.Do(context.Background(), "JSON.GET", keys.Kanji(id))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	resolvedVocab := &model.ResolvedVocabulary{
		ID:                vocab.ID,
		Object:            vocab.Object,
		Slug:              vocab.Slug,
		Characters:        vocab.Characters,
		Meanings:          vocab.Meanings,
		ReadingMnemonic:   vocab.ReadingMnemonic,
		Readings:          vocab.Readings,
		ContextSentences:  vocab.ContextSentences,
		ComponentSubjects: make([]model.Kanji, len(componentCmds)),
	}

	for i, cmd := range componentCmds {
		kanji := model.Kanji{}

		err = json.Unmarshal([]byte(cmd.Val().(string)), &kanji)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal json: %w", err)
		}

		resolvedVocab.ComponentSubjects[i] = kanji
	}

	return resolvedVocab, nil
}
