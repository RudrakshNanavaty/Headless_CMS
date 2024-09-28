package main

import (
	"github.com/gin-gonic/gin"
	"headless-cms/controllers/auth"
	"headless-cms/initializers"
	"headless-cms/middlewares"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/signup", auth.SignUp)
	r.POST("/login", auth.Login)
	r.POST("/logout", auth.Logout)
	r.GET("/auth", middlewares.RequireAuth, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Authenticated",
		})
	})

	r.Run()
}
