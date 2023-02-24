package keys

import "fmt"

func MeaningMnemonicIds() string {
	return "meaning_mnemonic_ids"
}

func MeaningMnemonic(id int) string {
	return fmt.Sprintf("meaning_mnemonic:%d", id)
}

func Kanji(id int) string {
	return fmt.Sprintf("kanji:%d", id)
}

func Radical(id int) string {
	return fmt.Sprintf("radical:%d", id)
}

func Vocabulary(id int) string {
	return fmt.Sprintf("vocabulary:%d", id)
}

func SearchIndex() string {
	return "search_index"
}
