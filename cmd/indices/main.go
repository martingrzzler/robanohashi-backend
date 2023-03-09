package main

import (
	"fmt"
	"log"
	"robanohashi/internal/config"
	"robanohashi/persist"
)

func main() {
	cfg := config.NewConfig()
	db, err := persist.Connect(cfg.RedisURL, cfg.RedisPassword)
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
