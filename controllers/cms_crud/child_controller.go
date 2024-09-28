package cms_crud

import (
	"github.com/gin-gonic/gin"
	"headless-cms/initializers"
	"headless-cms/types"
	"net/http"
)

func addChild(c *gin.Context) {
	var child types.Child

	err := c.BindJSON(&child)
	if err != nil {
		return
	}

	saved := initializers.DB.Create(&child)

	if saved.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Error saving child",
			"error_message": saved.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Child added successfully",
		"child":   child,
	})
}

func getChildren(c *gin.Context) {
	var children []types.Content
	id := c.Param("id")

	retrieved := initializers.DB.Preload("Children").First(&children, id)
	if retrieved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving children",
			"error_message": retrieved.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "Children retrieved successfully",
		"children": children,
	})
}

func getChild(c *gin.Context) {
	var child types.Child
	id := c.Param("id")

	retrieved := initializers.DB.First(&child, id)
	if retrieved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving child",
			"error_message": retrieved.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Child retrieved successfully",
		"child":   child,
	})
}

func updateChild(c *gin.Context) {
	var child types.Child
	err := c.BindJSON(&child)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}

	saved := initializers.DB.Model(&child).Where("parent_id = ? AND child_id = ?", child.ParentID, child.ChildID).Updates(&child)
	if saved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error updating child",
			"error_message": saved.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Child updated successfully",
		"child":   child,
	})
}

func deleteChild(c *gin.Context) {
	var child types.Child
	err := c.BindJSON(&child)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}

	deleted := initializers.DB.Where("parent_id = ? AND child_id = ?", child.ParentID, child.ChildID).Delete(&child)
	if deleted.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error deleting child",
			"error_message": deleted.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Child deleted successfully",
	})
}
