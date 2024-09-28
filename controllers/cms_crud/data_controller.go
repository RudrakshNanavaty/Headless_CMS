package cms_crud

import (
	"github.com/gin-gonic/gin"
	"headless-cms/initializers"
	"headless-cms/types"
	"net/http"
)

func addData(c *gin.Context) {
	var dataData types.Data
	err := c.BindJSON(&dataData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	// Save the data to the database
	saved := initializers.DB.Create(&dataData)
	if saved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error saving data",
			"error_message": saved.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data saved successfully",
		"data":    dataData,
	})
}

func getAllData(c *gin.Context) {
	var datasData []types.Data
	retrieved := initializers.DB.Find(&datasData)
	if retrieved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving data",
			"error_message": retrieved.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data retrieved successfully",
		"data":    datasData,
	})
}

func getData(c *gin.Context) {
	var dataData types.Data
	id := c.Param("id")
	retrieved := initializers.DB.First(&dataData, id)
	if retrieved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving data",
			"error_message": retrieved.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data retrieved successfully",
		"data":    dataData,
	})
}

func updateData(c *gin.Context) {
	var dataData types.Data
	err := c.BindJSON(&dataData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	// Save the data to the database
	saved := initializers.DB.Save(&dataData)
	if saved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error saving data",
			"error_message": saved.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data updated successfully",
		"data":    dataData,
	})
}

func deleteData(c *gin.Context) {
	var dataData types.Data
	id := c.Param("id")
	retrieved := initializers.DB.First(&dataData, id)
	if retrieved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving data",
			"error_message": retrieved.Error.Error(),
		})
		return
	}
	// Delete the data from the database
	deleted := initializers.DB.Delete(&dataData)
	if deleted.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error deleting data",
			"error_message": deleted.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data deleted successfully",
	})
}
