package persist

import (
	"context"
	"fmt"
	"log"
)

func ExampleDB_GetRadical() {
	radical, err := db.GetRadical(context.Background(), 1)
	if err != nil {
		log.Fatalf("failed to get radical: %v", err)
	}

	fmt.Println(radical)

	_, err = db.GetRadical(context.Background(), 2)

	fmt.Println(err)

	// Output:
	// &{1 radical one 一  [2] [{one true}] This is a radical for one}
	// subject is not a radical
}

func ExampleDB_GetRadicalResolved() {
	r, _ := db.GetRadical(context.Background(), 1)

	radical, err := db.GetRadicalResolved(context.Background(), r)
	if err != nil {
		log.Fatalf("failed to get radical: %v", err)
	}

	fmt.Println(radical)

	// Output:
	// &{1 radical one 一  [{2 kanji 一 one read it out loud [3] [{one true}] [{いち true onyomi}] [1] []}] [{one true}] This is a radical for one}
}
