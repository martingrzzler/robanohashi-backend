package persist

import (
	"context"
	"fmt"
	"log"
	"robanohashi/internal/dto"
	"robanohashi/internal/model"
)

func ExampleDB_UpdateMeaningMnemonic() {
	upd := dto.UpdateMeaningMnemonic{
		ID:   "1",
		Text: "updated meaning mnemonic",
	}
	err := db.UpdateMeaningMnemonic(context.Background(), upd)

	if err != nil {
		log.Fatalf("failed to update meaning mnemonic: %v", err)
	}

	mm1, _ := db.GetMeaningMnemonic(context.Background(), "1")

	fmt.Println(mm1.Text)

	err = seed()
	if err != nil {
		log.Fatalf("failed to reseed redis: %v", err)
	}

	// Output:
	// updated meaning mnemonic
}

func ExampleDB_DeleteMeaningMnemonic() {
	err := db.DeleteMeaningMnemonic(context.Background(), "1")

	if err != nil {
		log.Fatalf("failed to delete meaning mnemonic: %v", err)
	}

	_, err = db.GetMeaningMnemonic(context.Background(), "1")

	fmt.Println(err)

	err = seed()

	if err != nil {
		log.Fatalf("failed to reseed redis: %v", err)
	}

	// Output:
	// failed to get json: redis: nil
}

func ExampleDB_GetMeaningMnemonic() {
	mm1, _ := db.GetMeaningMnemonic(context.Background(), "1")

	fmt.Println(mm1.ID)

	// Output:
	// 1
}

func ExampleDB_UpvoteMeaningMnemonic() {
	status, err := db.UpvoteMeaningMnemonic(context.Background(), "1", "testuser")

	fmt.Println(status)

	if err != nil {
		log.Fatalf("failed to upvote meaning mnemonic: %v", err)
	}

	mm1, _ := db.GetMeaningMnemonic(context.Background(), "1")

	fmt.Println(mm1.VotingCount)

	// undo
	status, _ = db.UpvoteMeaningMnemonic(context.Background(), "1", "testuser")

	fmt.Println(status)

	mm1, _ = db.GetMeaningMnemonic(context.Background(), "1")

	fmt.Println(mm1.VotingCount)

	err = seed()

	if err != nil {
		log.Fatalf("failed to reseed redis: %v", err)
	}

	// Output:
	// upvoted
	// 1
	// removed upvote
	// 0
}

func ExampleDB_DownvoteMeaningMnemonic() {
	status, err := db.DownvoteMeaningMnemonic(context.Background(), "1", "testuser")

	fmt.Println(status)

	if err != nil {
		log.Fatalf("failed to downvote meaning mnemonic: %v", err)
	}

	mm1, _ := db.GetMeaningMnemonic(context.Background(), "1")

	fmt.Println(mm1.VotingCount)

	// undo
	status, _ = db.DownvoteMeaningMnemonic(context.Background(), "1", "testuser")

	fmt.Println(status)

	mm1, _ = db.GetMeaningMnemonic(context.Background(), "1")

	fmt.Println(mm1.VotingCount)

	err = seed()

	if err != nil {
		log.Fatalf("failed to reseed redis: %v", err)
	}

	// Output:
	// downvoted
	// -1
	// removed downvote
	// 0
}

func ExampleDB_ToggleFavoriteMeaningMnemonic() {
	status, err := db.ToggleFavoriteMeaningMnemonic(context.Background(), "1", "testuser")

	fmt.Println(status)

	if err != nil {
		log.Fatalf("failed to toggle favorite meaning mnemonic: %v", err)
	}

	mm, _ := db.GetMeaningMnemonic(context.Background(), "1")
	mms, _ := db.ResolveMeaningMnemonics(context.Background(), "testuser", []model.MeaningMnemonic{*mm})

	fmt.Println(mms[0].Favorite)

	// undo
	status, _ = db.ToggleFavoriteMeaningMnemonic(context.Background(), "1", "testuser")

	fmt.Println(status)

	mm, _ = db.GetMeaningMnemonic(context.Background(), "1")
	mms, _ = db.ResolveMeaningMnemonics(context.Background(), "testuser", []model.MeaningMnemonic{*mm})

	fmt.Println(mms[0].Favorite)

	err = seed()

	if err != nil {
		log.Fatalf("failed to reseed redis: %v", err)
	}

	// Output:
	// added
	// true
	// removed
	// false
}

func ExampleDB_GetMeaningMnemonicsBySubjectID() {
	mms, _ := db.GetMeaningMnemonicsBySubjectID(context.Background(), 2)

	fmt.Println(mms.TotalCount)
	fmt.Println(mms.Items[0].ID, mms.Items[1].ID)

	// Output:
	// 2
	// 2 1
}

func ExampleDB_GetMeaningMnemonicsByUser() {
	mms, _ := db.GetMeaningMnemonicsByUser(context.Background(), "testuser")

	fmt.Println(mms.TotalCount)
	fmt.Println(mms.Items[0].ID, mms.Items[1].ID)

	// Output:
	// 2
	// 2 1
}

func ExampleDB_GetFavoriteMeaningMnemonics() {
	_, _ = db.ToggleFavoriteMeaningMnemonic(context.Background(), "1", "testuser")
	mms, _ := db.GetFavoriteMeaningMnemonics(context.Background(), "testuser")

	fmt.Println(len(mms))
	fmt.Println(mms[0].ID)

	// Output:
	// 1
	// 1
}

func ExampleDB_ResolveMeaningMnemonics() {
	mms, _ := db.GetMeaningMnemonicsByUser(context.Background(), "testuser")

	resolved, err := db.ResolveMeaningMnemonics(context.Background(), "testuser", mms.Items)

	if err != nil {
		log.Fatalf("failed to resolve meaning mnemonics: %v", err)
	}

	fmt.Println(len(resolved))
	fmt.Println(resolved[0].ID, resolved[1].ID)
	fmt.Println(resolved[0].Favorite, resolved[0].Me, resolved[0].Upvoted, resolved[0].Downvoted)

	// Output:
	// 2
	// 2 1
	// false true false false
}
