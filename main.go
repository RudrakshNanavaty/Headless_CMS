package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"headless-cms/utils"
)

func main() {

	utils.LoadEnv()

	// Loading environment variables
	DB_HOST := utils.GetEnv("DB_HOST")
	DB_USER := utils.GetEnv("DB_USER")
	DB_PASSWORD := utils.GetEnv("DB_PASSWORD")
	DB_PORT := utils.GetEnv("DB_PORT")
	DB_NAME := utils.GetEnv("DB_NAME")
	DB_SSLMODE := utils.GetEnv("DB_SSLMODE")

	// Connecting to DB
	_, err := utils.ConnectDBByCredentials(DB_HOST, DB_USER, DB_PASSWORD, DB_PORT, DB_NAME, DB_SSLMODE)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to DB")
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World",
			})
		})
		if err := r.Run(":3000"); err != nil { // listen and serve on 0.0.0.0:3000
			fmt.Println(err)
		}
	}
}
