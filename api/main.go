package main

import (
	"log"
	"net/http"
	"robanohashi/persist"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := persist.Connect()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	router := gin.Default()
	router.GET("/search/:query", func(c *gin.Context) {
		query := c.Param("query")

		res, err := db.SearchSubjects(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(200, gin.H{
			"data": res,
		})
	})
	router.Run(":5000")
}
