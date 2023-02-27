package keys

import "fmt"

// single int that is incremented every time a new meaning mnemonic is added
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

func SubjectIndex() string {
	return "subject_index"
}

func MeaningMnemonicIndex() string {
	return "meaning_mnemonic_index"
}
