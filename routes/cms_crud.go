package routes

import (
	"github.com/gin-gonic/gin"
	"headless-cms/controllers/cms_crud"
	"headless-cms/middlewares"
)

func LoadCMSCRUDRoutes(R *gin.RouterGroup) {
	// Attribute
	R.GET("/attributes", cms_crud.GetAttributes)
	R.POST("/attribute", middlewares.RequireAdmin, cms_crud.AddAttribute)
	R.GET("/attribute/:id", cms_crud.GetAttribute)
	R.PUT("/attribute/:id", middlewares.RequireAdmin, cms_crud.UpdateAttribute)
	R.DELETE("/attribute/:id", middlewares.RequireAdmin, cms_crud.DeleteAttribute)

	// Content
	R.GET("/contents", cms_crud.GetContents)
	R.POST("/content", middlewares.RequireAdmin, cms_crud.AddContent)
	R.GET("/content/:id", cms_crud.GetContent)
	R.PUT("/content/:id", middlewares.RequireAdmin, cms_crud.UpdateContent)
	R.DELETE("/content/:id", middlewares.RequireAdmin, cms_crud.DeleteContent)

	// Data
	R.GET("/alldata", cms_crud.GetAllData)
	R.POST("/data", middlewares.RequireAdmin, cms_crud.AddData)
	R.GET("/data/:id", cms_crud.GetData)
	R.PUT("/data/:id", middlewares.RequireAdmin, cms_crud.UpdateData)
	R.DELETE("/data/:id", middlewares.RequireAdmin, cms_crud.DeleteData)
	R.PATCH("/data/:id", middlewares.RequireAdmin, cms_crud.UploadFile)

	// Type
	R.GET("/types", cms_crud.GetTypes)
	R.POST("/type", middlewares.RequireAdmin, cms_crud.AddType)
	R.GET("/type/:id", cms_crud.GetType)
	R.PUT("/type/:id", middlewares.RequireAdmin, cms_crud.UpdateType)

	// Child
	R.GET("/children", cms_crud.GetChildren)
	R.POST("/child", middlewares.RequireAdmin, cms_crud.AddChild)
	R.GET("/child/:id", cms_crud.GetChild)
	R.PUT("/child/:id", middlewares.RequireAdmin, cms_crud.UpdateChild)
	R.DELETE("/child/:id", middlewares.RequireAdmin, cms_crud.DeleteChild)
}
