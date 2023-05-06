package persist

import (
	"context"
	"fmt"
	"robanohashi/persist/keys"
)

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
		"$.user_id", "AS", "user_id", "TAG",
		"$.voting_count", "AS", "voting_count", "NUMERIC", "SORTABLE",
	).Err()

	return err
}

func (db *DB) createSubjectIndex() error {
	// Create indices
	err := db.rdb.Do(context.Background(),
		"FT.CREATE",
		keys.SubjectIndex(),
		"ON", "JSON",
		"PREFIX", "3", "radical:", "kanji:", "vocabulary:",
		"SCHEMA",
		"$.characters", "AS", "characters", "TAG",
		"$.meanings.*.meaning", "AS", "meaning", "TEXT",
		"$.readings.*.reading", "AS", "reading", "TAG",
		"$.readings.*.romaji", "AS", "romaji", "TAG",
		"$.source", "AS", "source", "NUMERIC", "SORTABLE",
	).Err()

	return err
}
