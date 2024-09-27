package cms_crud

import (
	"github.com/gin-gonic/gin"
	"headless-cms/initializers"
	"headless-cms/types"
)

func addType(c *gin.Context) {
	var typeData types.Type
	err := c.BindJSON(&typeData)
	if err != nil {
		c.JSON(400, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	// Save the type to the database
	saved := initializers.DB.Create(&typeData)

	if saved.Error != nil {
		c.JSON(500, gin.H{
			"message":       "Error saving type",
			"error_message": saved.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Type saved successfully",
		"type":    typeData,
	})
}

func getTypes(c *gin.Context) {
	var typesData []types.Type
	retrieved := initializers.DB.Find(&typesData)
	if retrieved.Error != nil {
		c.JSON(500, gin.H{
			"message":       "Error retrieving types",
			"error_message": retrieved.Error.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "Types retrieved successfully",
		"types":   typesData,
	})
}

func getType(c *gin.Context) {
	var typeData types.Type
	id := c.Param("id")
	retrieved := initializers.DB.First(&typeData, id)
	if retrieved.Error != nil {
		c.JSON(500, gin.H{
			"message":       "Error retrieving type",
			"error_message": retrieved.Error.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "Type retrieved successfully",
		"type":    typeData,
	})
}

func updateType(c *gin.Context) {
	var typeData types.Type
	id := c.Param("id")
	err := c.BindJSON(&typeData)
	if err != nil {
		c.JSON(400, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}

	saved := initializers.DB.Model(&typeData).Where("id = ?", id).Updates(&typeData)

	if saved.Error != nil {
		c.JSON(500, gin.H{
			"message":       "Error updating type",
			"error_message": saved.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Type updated successfully",
		"type":    typeData,
	})
}
