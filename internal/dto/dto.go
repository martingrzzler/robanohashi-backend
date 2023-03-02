package dto

import (
	"encoding/json"
	"errors"
	"robanohashi/internal/model"
)

type RawUnmarshaler[T any] interface {
	UnmarshalRaw(data any) (T, error)
}

func ParseFTSearchResult[T RawUnmarshaler[T]](result any) (int64, []T, error) {
	items := make([]T, 0)

	for i, item := range result.([]any)[1:] {
		if i%2 == 0 {
			continue
		}

		obj := *new(T)

		parsed, err := obj.UnmarshalRaw(item.([]any)[1])
		if err != nil {
			return 0, nil, err
		}

		items = append(items, parsed)
	}

	totalCount := result.([]interface{})[0].(int64)

	return totalCount, items, nil
}

type MeaningMnemonic struct {
	model.MeaningMnemonic
}

func (m MeaningMnemonic) UnmarshalRaw(data any) (MeaningMnemonic, error) {
	s, ok := data.(string)
	if !ok {
		return MeaningMnemonic{}, errors.New("could not convert data to string")
	}

	mm := MeaningMnemonic{}

	err := json.Unmarshal([]byte(s), &mm)
	if err != nil {
		return MeaningMnemonic{}, err
	}
	return mm, nil
}