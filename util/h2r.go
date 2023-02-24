package util

import "strings"

// source https://www.lexilogos.com/keyboard/hiragana_conversion.htm
var h2r map[string]string = map[string]string{
	"ん": "n", "きゃ": "kya", "きゅ": "kyu", "きょ": "kyo",
	"にゃ": "nya", "にゅ": "nyu", "にょ": "nyo", "しゃ": "sha",
	"しゅ": "shu", "しょ": "sho", "ちゃ": "cha", "ちゅ": "chu",
	"ちょ": "cho", "ひゃ": "hya", "ひゅ": "hyu", "ひょ": "hyo",
	"みゃ": "mya", "みゅ": "myu", "みょ": "myo", "りゃ": "rya",
	"りゅ": "ryu", "りょ": "ryo", "ぎゃ": "gya", "ぎゅ": "gyu",
	"ぎょ": "gyo", "びゃ": "bya", "びゅ": "byu", "びょ": "byo",
	"ぴゃ": "pya", "ぴゅ": "pyu", "ぴょ": "pyo", "じゃ": "ja",
	"じゅ": "ju", "じょ": "jo", "し": "shi", "ち": "chi", "つ": "tsu",
	"ば": "ba", "だ": "da", "が": "ga", "は": "ha", "か": "ka",
	"ま": "ma", "な": "na", "ぱ": "pa", "ら": "ra", "さ": "sa",
	"た": "ta", "わ": "wa", "や": "ya", "ざ": "za", "あ": "a",
	"べ": "be", "で": "de", "げ": "ge", "へ": "he", "け": "ke",
	"め": "me", "ね": "ne", "ぺ": "pe", "れ": "re", "せ": "se",
	"て": "te", "ゑ": "we", "ぜ": "ze", "え": "e", "び": "bi",
	"ぎ": "gi", "ひ": "hi", "き": "ki", "み": "mi", "に": "ni",
	"ぴ": "pi", "り": "ri", "ゐ": "wi", "じ": "ji", "い": "i",
	"ぼ": "bo", "ど": "do", "ご": "go", "ほ": "ho", "こ": "ko",
	"も": "mo", "の": "no", "ぽ": "po", "ろ": "ro", "そ": "so",
	"と": "to", "を": "wo", "よ": "yo", "ぞ": "zo", "お": "o",
}

func HiraganaToRomaji(hiragana string) string {
	var romaji strings.Builder
	var prev rune

	for _, char := range hiragana {
		if translit, ok := h2r[string(char)]; ok {
			// If the previous character was 'n' and this one is a vowel, add a macron to indicate a long vowel.
			if prev == 'ん' && strings.ContainsAny(translit, "aiueo") {
				romaji.WriteString("¯")
			}
			romaji.WriteString(translit)
			prev = char
		} else {
			// If the character is not in the mapping table, just copy it to the output string as-is.
			romaji.WriteRune(char)
			prev = char
		}
	}

	return romaji.String()

}
