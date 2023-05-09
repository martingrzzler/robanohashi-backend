package persist

import (
	"context"
	"fmt"
	"log"
	"robanohashi/persist/keys"
)

func ExampleDB_ToggleSubjectBookmarked() {
	status, err := db.ToggleSubjectBookmarked(context.Background(), keys.Kanji(2), "testuser")

	if err != nil {
		log.Fatalf("failed to toggle subject bookmark: %v", err)
	}

	fmt.Println(status)

	// undo
	status, _ = db.ToggleSubjectBookmarked(context.Background(), keys.Kanji(2), "testuser")

	fmt.Println(status)

	err = seed()

	if err != nil {
		log.Fatalf("failed to reseed redis: %v", err)
	}

	// Output:
	// added
	// removed
}

func ExampleDB_SubjectBookmarked() {
	db.ToggleSubjectBookmarked(context.Background(), keys.Radical(1), "testuser")
	db.ToggleSubjectBookmarked(context.Background(), keys.Kanji(2), "testuser")

	bookmarked, err := db.SubjectBookmarked(context.Background(), keys.Radical(1), "testuser")

	if err != nil {
		log.Fatalf("failed to get subject bookmark status: %v", err)
	}

	fmt.Println(bookmarked)

	bookmarked, _ = db.SubjectBookmarked(context.Background(), keys.Kanji(2), "testuser")

	fmt.Println(bookmarked)

	bookmarked, _ = db.SubjectBookmarked(context.Background(), keys.Vocabulary(3), "testuser")

	fmt.Println(bookmarked)

	err = seed()

	if err != nil {
		log.Fatalf("failed to reseed redis: %v", err)
	}

	// Output:
	// true
	// true
	// false
}

func ExampleDB_GetUserBookmarkedSubjects() {

	db.ToggleSubjectBookmarked(context.Background(), keys.Radical(1), "testuser")
	db.ToggleSubjectBookmarked(context.Background(), keys.Kanji(2), "testuser")
	db.ToggleSubjectBookmarked(context.Background(), keys.Vocabulary(3), "testuser")

	subjects, err := db.GetUserBookmarkedSubjects(context.Background(), "testuser")
	if err != nil {
		log.Fatalf("failed to get user bookmarked subjects: %v", err)
	}

	fmt.Println(subjects)

	// remove radical
	db.ToggleSubjectBookmarked(context.Background(), keys.Radical(1), "testuser")

	subjects, _ = db.GetUserBookmarkedSubjects(context.Background(), "testuser")

	fmt.Println(subjects)

	err = seed()

	if err != nil {
		log.Fatalf("failed to reseed redis: %v", err)
	}

	// Output:
	// &{0 [{2 kanji one wanikani 一  [いち] [one]} {3 vocabulary one dictionary 一  [いち] [one]} {1 radical one wanikani 一  [] [one]}]}
	// &{0 [{2 kanji one wanikani 一  [いち] [one]} {3 vocabulary one dictionary 一  [いち] [one]}]}
}
