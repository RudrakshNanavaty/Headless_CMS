package super_admin

import (
	"github.com/gin-gonic/gin"
	"headless-cms/config/roles"
	"headless-cms/initializers"
	"headless-cms/types"
	"net/http"
)

func RegisterAdmin(c *gin.Context) {
	// Register admin
	var admin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.BindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	var newAdmin types.User
	newAdmin.Username = admin.Username
	newAdmin.Password = admin.Password
	newAdmin.RoleType = roles.Admin
	// Save the admin to the database
	saved := initializers.DB.Create(&newAdmin)
	if saved.Error != nil {
		if saved.Error.Error() == "UNIQUE constraint failed: users.username" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Username already exists",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error saving admin",
			"error_message": saved.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin saved successfully",
	})
}

func DeleteAdmin(c *gin.Context) {
	// Delete admin
	var admin types.User
	id := c.Param("id")
	deleted := initializers.DB.Delete(&admin, id)
	if deleted.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error deleting admin",
			"error_message": deleted.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Admin deleted successfully",
	})
}

func UpdateAdmin(c *gin.Context) {
	// Update admin
	var adminData types.User
	err := c.BindJSON(&adminData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	id := c.Param("id")
	// Update the admin in the database
	updated := initializers.DB.Model(&adminData).Where("id = ?", id).Updates(&adminData)
	if updated.Error != nil {
		if updated.Error.Error() == "UNIQUE constraint failed: users.username" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Username already exists",
			})
			return
		}
		// handle not found error
		if updated.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Admin not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error updating admin",
			"error_message": updated.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Admin updated successfully",
	})
}

func GetAccount(c *gin.Context) {
	// Get account
	var account types.User
	id := c.Param("id")
	retrieved := initializers.DB.Select("id", "username", "role_type").First(&account, id)
	if retrieved.Error != nil {
		if retrieved.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Account not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving admin",
			"error_message": retrieved.Error.Error(),
			"data":          retrieved,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin retrieved successfully",
		"account": account,
	})
}

func GetAllUsers(c *gin.Context) {
	// Get all users
	var users []types.User
	retrieved := initializers.DB.Select("id", "username", "role_type").Find(&users, "role_type = ?", roles.User)
	if retrieved.Error != nil {
		if retrieved.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No users found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving admins",
			"error_message": retrieved.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Admins retrieved successfully",
		"users":   users,
	})
}

func GetAllAdmins(c *gin.Context) {
	// Get all admins
	var admins []types.User
	retrieved := initializers.DB.Select("id", "username", "role_type").Find(&admins, "role_type = ?", roles.Admin)
	if retrieved.Error != nil {
		if retrieved.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No admins found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving admins",
			"error_message": retrieved.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Admins retrieved successfully",
		"admins":  admins,
	})
}

func PromoteUserToAdmin(c *gin.Context) {
	var id = c.Param("id")
	var user types.User
	retrieved := initializers.DB.First(&user, id)
	if retrieved.Error != nil {
		if retrieved.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	user.RoleType = roles.Admin
	updated := initializers.DB.Save(&user)
	if updated.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error promoting user to admin",
			"error_message": updated.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User promoted to admin successfully",
	})
}

func DemoteAdminToUser(c *gin.Context) {
	var id = c.Param("id")
	var user types.User
	retrieved := initializers.DB.First(&user, id)
	if retrieved.Error != nil {
		if retrieved.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	user.RoleType = roles.User
	updated := initializers.DB.Save(&user)
	if updated.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error demoting admin to user",
			"error_message": updated.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Admin demoted to user successfully",
	})
}
