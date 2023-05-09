package keys

import "fmt"

// single int that is incremented every time a new meaning mnemonic is added
func MeaningMnemonicIds() string {
	return "meaning_mnemonic_ids"
}

func MeaningMnemonic(id string) string {
	return fmt.Sprintf("meaning_mnemonic:%s", id)
}

func MeaningMnemonicDownVoters(id string) string {
	return fmt.Sprintf("meaning_mnemonic_down_voters:%s", id)
}

func MeaningMnemonicUpVoters(id string) string {
	return fmt.Sprintf("meaning_mnemonic_up_voters:%s", id)
}

func MeaningMnemonicFavorites(userId string) string {
	return fmt.Sprintf("meaning_mnemonic_favorites:%s", userId)
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

func UserBoomarks(userId string) string {
	return fmt.Sprintf("user_bookmarks:%s", userId)
}
