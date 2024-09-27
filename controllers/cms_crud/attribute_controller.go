package cms_crud

import (
	"github.com/gin-gonic/gin"
	"headless-cms/initializers"
	"headless-cms/types"
)

func addAttribute(c *gin.Context) {
	var attributeData types.Attribute
	err := c.BindJSON(&attributeData)
	if err != nil {
		c.JSON(400, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	// Save the attribute to the database
	saved := initializers.DB.Create(&attributeData)
	if saved.Error != nil {
		c.JSON(500, gin.H{
			"message":       "Error saving attribute",
			"error_message": saved.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":   "Attribute saved successfully",
		"attribute": attributeData,
	})
}

func getAttributes(c *gin.Context) {
	var attributesData []types.Attribute
	retrieved := initializers.DB.Find(&attributesData)
	if retrieved.Error != nil {
		c.JSON(500, gin.H{
			"message":       "Error retrieving attributes",
			"error_message": retrieved.Error.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message":    "Attributes retrieved successfully",
		"attributes": attributesData,
	})
}

func getAttribute(c *gin.Context) {
	var attributeData types.Attribute
	id := c.Param("id")
	retrieved := initializers.DB.First(&attributeData, id)
	if retrieved.Error != nil {
		c.JSON(500, gin.H{
			"message":       "Error retrieving attribute",
			"error_message": retrieved.Error.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message":   "Attribute retrieved successfully",
		"attribute": attributeData,
	})
}

func updateAttribute(c *gin.Context) {
	var attributeData types.Attribute
	err := c.BindJSON(&attributeData)
	if err != nil {
		c.JSON(400, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	id := c.Param("id")
	// Update the attribute in the database
	updated := initializers.DB.Model(&attributeData).Where("id = ?", id).Updates(attributeData)
	if updated.Error != nil {
		c.JSON(500, gin.H{
			"message":       "Error updating attribute",
			"error_message": updated.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":   "Attribute updated successfully",
		"attribute": attributeData,
	})
}

func deleteAttribute(c *gin.Context) {
	var attributeData types.Attribute
	id := c.Param("id")
	// Delete the attribute from the database
	deleted := initializers.DB.Where("id = ?", id).Delete(&attributeData)
	if deleted.Error != nil {
		c.JSON(500, gin.H{
			"message":       "Error deleting attribute",
			"error_message": deleted.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Attribute deleted successfully",
	})
}
