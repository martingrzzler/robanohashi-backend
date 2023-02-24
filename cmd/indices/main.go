package main

import (
	"context"
	"fmt"
	"log"
	"robanohashi/db"
	"robanohashi/db/keys"
)

func main() {
	rdb, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer rdb.Close()

	res, err := rdb.Do(context.Background(), "FT._LIST").Result()
	if err != nil {
		log.Fatal(err)
	}

	for _, index := range res.([]interface{}) {
		err := rdb.Do(context.Background(), "FT.DROPINDEX", index).Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create indices
	err = rdb.Do(context.Background(),
		"FT.CREATE",
		keys.SearchIndex(),
		"ON", "JSON",
		"PREFIX", "3", "kanji:", "radical:", "vocabulary:",
		"SCHEMA",
		"$.characters", "AS", "characters", "TAG",
		"$.meanings.*.meaning", "AS", "meaning", "TAG",
		"$.readings.*.reading", "AS", "reading", "TAG",
	).Err()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Indices created successfully")
}
