package cms_crud

import (
	"github.com/gin-gonic/gin"
	"headless-cms/initializers"
	"headless-cms/types"
	"net/http"
)

func addContent(c *gin.Context) {
	var contentData types.Content
	err := c.BindJSON(&contentData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	// Save the content to the database
	saved := initializers.DB.Create(&contentData)

	if saved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error saving content",
			"error_message": saved.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Content saved successfully",
		"content": contentData,
	})
}

func getContents(c *gin.Context) {
	var contentsData []types.Content
	retrieved := initializers.DB.Find(&contentsData)
	if retrieved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving contents",
			"error_message": retrieved.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "Contents retrieved successfully",
		"contents": contentsData,
	})
}

func getContent(c *gin.Context) {
	var contentData types.Content
	id := c.Param("id")
	retrieved := initializers.DB.First(&contentData, id)
	if retrieved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving content",
			"error_message": retrieved.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Content retrieved successfully",
		"content": contentData,
	})
}

func updateContent(c *gin.Context) {
	var contentData types.Content
	err := c.BindJSON(&contentData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	// Save the content to the database
	saved := initializers.DB.Save(&contentData).Where("type_id = ?", contentData.TypeID)

	if saved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error saving content",
			"error_message": saved.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Content updated successfully",
		"content": contentData,
	})
}

func deleteContent(c *gin.Context) {
	var contentData types.Content
	err := c.BindJSON(&contentData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	deleted := initializers.DB.Delete(&contentData)
	if deleted.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error deleting content",
			"error_message": deleted.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Content deleted successfully",
	})
}
