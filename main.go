package main

import (
	"github.com/gin-gonic/gin"
	"headless-cms/initializers"
	"headless-cms/middlewares"
	"headless-cms/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.CreateSuperAdminIfNotExists()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	superAdminRoutes := r.Group("/super-admin", middlewares.RequireSuperAdmin)
	routes.LoadSuperAdminRoutes(superAdminRoutes)

	authRoutes := r.Group("/auth")
	routes.LoadAuthRoutes(authRoutes)

	r.Run()
}
