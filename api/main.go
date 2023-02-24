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
	router.Run(":5000")
}
