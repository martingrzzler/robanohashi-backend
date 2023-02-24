package main

import (
	"fmt"
	"log"
	"robanohashi/db"
)

func main() {
	db, err := db.Connect()
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
