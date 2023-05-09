package persist

import (
	"context"
	"fmt"
	"robanohashi/internal/dto"
	"robanohashi/persist/keys"

	"github.com/redis/go-redis/v9"
)

func (db *DB) ToggleSubjectBookmarked(ctx context.Context, subjectKey string, uid string) (string, error) {
	return db.toggleSetValue(ctx, keys.UserBoomarks(uid), subjectKey)
}

func (db *DB) GetUserBookmarkedSubjects(ctx context.Context, uid string) (*dto.List[dto.SubjectPreview], error) {
	sKeys, err := db.rdb.SMembers(ctx, keys.UserBoomarks(uid)).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to get user bookmarked subjects: %w", err)
	}

	pipe := db.rdb.Pipeline()
	sKeysCmds := make([]*redis.Cmd, len(sKeys))

	for i, key := range sKeys {
		sKeysCmds[i] = pipe.Do(ctx, "JSON.GET", key)
	}

	_, err = pipe.Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to execute pipeline: %w", err)
	}

	subjects := make([]dto.SubjectPreview, len(sKeysCmds))

	for i, cmd := range sKeysCmds {
		s, err := dto.SubjectPreview{}.UnmarshalRaw(cmd.Val().(string))

		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal subject: %w", err)
		}

		subjects[i] = s
	}

	return &dto.List[dto.SubjectPreview]{Items: subjects}, nil
}
