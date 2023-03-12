package persist

import (
	"context"
	"fmt"
	"log"
)

func ExampleDB_GetKanji() {
	kanji, err := db.GetKanji(context.Background(), 2)
	if err != nil {
		log.Fatalf("failed to get kanji: %v", err)
	}

	fmt.Println(kanji)

	_, err = db.GetKanji(context.Background(), 1)

	fmt.Println(err)

	// Output:
	// &{2 kanji 一 one read it out loud [3] [{one true}] [{いち true onyomi}] [1] []}
	// subject is not a kanji
}

func ExampleDB_GetKanjiResolved() {
	k, _ := db.GetKanji(context.Background(), 2)

	kanji, err := db.GetKanjiResolved(context.Background(), k)
	if err != nil {
		log.Fatalf("failed to get kanji: %v", err)
	}

	fmt.Println(kanji)

	// Output:
	// &{2 kanji 一 one read it out loud [{3 vocabulary one 一  [いち] [one]}] [{one true}] [{いち true onyomi}] [{1 radical one 一  [] [one]}] []}
}
