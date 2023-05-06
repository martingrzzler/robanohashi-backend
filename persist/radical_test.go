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
	// &{1 radical wanikani one 一  [2] [{one true}] This is a radical for one}
	// failed to get json: redis: nil
}

func ExampleDB_GetRadicalResolved() {
	r, _ := db.GetRadical(context.Background(), 1)

	radical, err := db.GetRadicalResolved(context.Background(), r)
	if err != nil {
		log.Fatalf("failed to get radical: %v", err)
	}

	fmt.Println(radical)

	// Output:
	// &{1 radical one wanikani 一  [{2 kanji one wanikani 一  [いち] [one]}] [{one true}] This is a radical for one}
}
