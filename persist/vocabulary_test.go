package persist

import (
	"context"
	"fmt"
	"log"
)

func ExampleDB_GetVocabulary() {
	v, err := db.GetVocabulary(context.Background(), 3)
	if err != nil {
		log.Fatalf("failed to get vocabulary: %v", err)
	}

	fmt.Println(v)

	_, err = db.GetVocabulary(context.Background(), 1)

	fmt.Println(err)

	// Output:
	// &{3 vocabulary dictionary one 一 [2] [{one true}] read it out loud [{I took one step forward 私は一歩前に進んだ わたしはいっぽうまえにすすんだ}] [{いち true ichi}]}
	// failed to get json: redis: nil
}

func ExampleDB_GetVocabularyResolved() {
	v, err := db.GetVocabulary(context.Background(), 3)
	if err != nil {
		log.Fatalf("failed to get vocabulary: %v", err)
	}

	vocab, err := db.GetVocabularyResolved(context.Background(), v)
	if err != nil {
		log.Fatalf("failed to get vocabulary: %v", err)
	}

	fmt.Println(vocab)

	// Output:
	// &{3 vocabulary one dictionary 一 [{2 kanji one wanikani 一  [いち] [one]}] [{one true}] read it out loud [{I took one step forward 私は一歩前に進んだ わたしはいっぽうまえにすすんだ}] [{いち true ichi}]}

}
