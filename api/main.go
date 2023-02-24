package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/search/:query", func(c *gin.Context) {
		query := c.Param("query")

		c.JSON(200, gin.H{
			"message": query,
		})
	})
	router.Run(":5000")
}
