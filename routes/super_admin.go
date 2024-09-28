package routes

import (
	"github.com/gin-gonic/gin"
	"headless-cms/controllers/super_admin"
)

// LoadSuperAdminRoutes loads the super admin routes
func LoadSuperAdminRoutes(R *gin.RouterGroup) {
	R.POST("/register-admin", super_admin.RegisterAdmin)
	R.GET("/list/admins", super_admin.GetAllAdmins)
	R.GET("/list/users", super_admin.GetAllUsers)
	R.GET("/account/:id", super_admin.GetAccount)
	R.PUT("/update/:id", super_admin.UpdateAdmin)
	R.PATCH("/promote/:id", super_admin.PromoteUserToAdmin)
	R.PATCH("/demote/:id", super_admin.DemoteAdminToUser)
	R.DELETE("/delete/:id", super_admin.DeleteAdmin)
}
