package util

import (
	"unicode"
	"unicode/utf8"
)

func SingleKanji(s string) bool {
	if len(s) == 3 && utf8.RuneCountInString(s) == 1 {
		r, _ := utf8.DecodeRuneInString(s)
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}

	return false
}
