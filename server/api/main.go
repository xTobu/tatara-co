package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xTobu/tatara-co/server/api/models"
	"github.com/xTobu/tatara-co/server/api/packages/helpers"
	"github.com/xTobu/tatara-co/server/api/packages/setting"
)

var ENV string

func main() {
	os.Setenv("ENV", ENV)
	fmt.Println(ENV)

	setting.Setup()
	models.Setup()

	router := gin.Default()

	api := router.Group("/api")
	api.Use()
	{
		api.GET("/test", func(c *gin.Context) {
			users, err := models.GetUsers()
			if err != nil {
				log.Fatalf("models.Setup err: %v", err)
			}
			helpers.OutputGormResult(users)
			c.JSON(200, gin.H{
				"status": "success",
				"data":   users,
			})
		})
		api.GET("/test/2", func(c *gin.Context) {

			c.JSON(200, gin.H{
				"status": "success",
			})
		})
	}

	router.Run(":3000")
}
