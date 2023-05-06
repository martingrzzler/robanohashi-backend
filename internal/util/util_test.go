package util

import (
	"testing"
)

func TestSingleKanji(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"水", true},
		{"犬", true},
		{"漢字", false},
		{"cat", false},
		{"", false},
		{"😀", false},
		{"あ", false},
		{"ア", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := SingleKanji(tc.input)

			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
