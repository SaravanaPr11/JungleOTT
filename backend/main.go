package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Health Check Route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run(":8080") // Start server on port 8080
}
