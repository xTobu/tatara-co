package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Simple group: v1
	api := router.Group("/api")
	{
		api.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "success",
				"data":   "Hello, world!",
			})
		})
	}
	router.Run(":3000")
}
