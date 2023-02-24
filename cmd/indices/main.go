package main

import (
	"fmt"
	"log"
	"robanohashi/persist"
)

func main() {
	db, err := persist.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.CreateIndices()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Indices created successfully")
}
