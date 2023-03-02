package main

import (
	"log"
	"robanohashi/api/controllers"
	"robanohashi/persist"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := persist.Connect()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/search", controllers.Search)
	router.GET("/kanji/:id", controllers.GetKanji)
	router.GET("/subject/:id/meaning_mnemonics", controllers.GetMeaningMnemonics)
	router.GET("/radical/:id", controllers.GetRadical)
	router.GET("/vocabulary/:id", controllers.GetVocabulary)
	router.Run(":5000")
}
