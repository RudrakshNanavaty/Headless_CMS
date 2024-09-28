package routes

import (
	"github.com/gin-gonic/gin"
	"headless-cms/controllers/auth"
	"headless-cms/middlewares"
)

func LoadAuthRoutes(c *gin.RouterGroup) {
	c.POST("/signup", auth.SignUp)
	c.POST("/login", auth.Login)
	c.POST("/logout", auth.Logout)
	c.GET("/auth", middlewares.RequireAuth, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Authenticated",
		})
	})
}
