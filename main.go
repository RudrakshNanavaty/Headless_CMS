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
	R := gin.Default()

	R.RedirectTrailingSlash = false

	r := R.Group("/api/v1")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	superAdminRoutes := r.Group("/super-admin", middlewares.RequireSuperAdmin)
	routes.LoadSuperAdminRoutes(superAdminRoutes)

	authRoutes := r.Group("/auth")
	routes.LoadAuthRoutes(authRoutes)

	cmsRoutes := r.Group("/cms", middlewares.RequireAuth)
	routes.LoadCMSCRUDRoutes(cmsRoutes)

	//R.PATCH("/upload", cms_crud.UploadFile)

	err := R.Run()
	if err != nil {
		panic(err)
	}
}
