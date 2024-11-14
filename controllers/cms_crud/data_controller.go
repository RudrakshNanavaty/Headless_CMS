package cms_crud

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"headless-cms/initializers"
	"headless-cms/types"
	"headless-cms/utils"
	"net/http"
	"path/filepath"
	"strings"
)

func AddData(c *gin.Context) {
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

	// update contentID with dataID
	var contentData types.Content
	retrieved := initializers.DB.First(&contentData, dataData.ContentID)
	if retrieved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving content",
			"error_message": retrieved.Error.Error(),
		})
		return
	}
	contentData.DataID = dataData.ID
	updated := initializers.DB.Model(&contentData).Where("id = ?", contentData.ID).Updates(&contentData)
	if updated.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error updating content",
			"error_message": updated.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data saved successfully",
		"data":    dataData,
	})
}

func GetAllData(c *gin.Context) {
	var datasData []types.Data
	retrieved := initializers.DB.Find(&datasData)
	if retrieved.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error retrieving data",
			"error_message": retrieved.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data retrieved successfully",
			"data":    datasData,
		})
	}
}

func GetData(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Data retrieved successfully",
		"data":    dataData,
	})
}

func UpdateData(c *gin.Context) {
	var dataData types.Data
	id := c.Param("id")
	err := c.BindJSON(&dataData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	// Save the data to the database
	//saved := initializers.DB.Save(&dataData)
	saved := initializers.DB.Model(&dataData).Where("id = ?", id).Updates(&dataData)
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

func DeleteData(c *gin.Context) {
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

func UploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	var errors []string
	var uploadedTextURLs, uploadedImageURLs, uploadedPdfURLs, uploadedCodeURLs []string

	files := form.File["files"]

	for _, file := range files {
		fileHeader := file

		f, err := fileHeader.Open()
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error opening file %s: %s", fileHeader.Filename, err.Error()))
			continue
		}
		defer f.Close()

		uploadedURL, err := utils.SaveFile(f, fileHeader)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error saving file %s: %s", fileHeader.Filename, err.Error()))
		} else {
			ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
			switch ext {
			case ".txt":
				uploadedTextURLs = append(uploadedTextURLs, uploadedURL)
			case ".jpg", ".jpeg", ".png", ".gif":
				uploadedImageURLs = append(uploadedImageURLs, uploadedURL)
			case ".pdf":
				uploadedPdfURLs = append(uploadedPdfURLs, uploadedURL)
			case ".go", ".py", ".js", ".java", "cpp", "c":
				uploadedCodeURLs = append(uploadedCodeURLs, uploadedURL)
			default:
				errors = append(errors, fmt.Sprintf("Unsupported file type %s", fileHeader.Filename))
			}
		}
	}
	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors})
	} else {
		// Save the uploaded URLs to the database
		var dataUpdated types.Data
		retrieved := initializers.DB.First(&dataUpdated, id)
		if retrieved.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message":       "Error retrieving data",
				"error_message": retrieved.Error.Error(),
			})
			return
		}

		dataUpdated.Text = append(dataUpdated.Text, uploadedTextURLs...)
		dataUpdated.ImageUrls = append(dataUpdated.ImageUrls, uploadedImageURLs...)
		dataUpdated.PdfUrls = append(dataUpdated.PdfUrls, uploadedPdfURLs...)
		dataUpdated.Code = append(dataUpdated.Code, uploadedCodeURLs...)

		updated := initializers.DB.Model(&dataUpdated).Where("id = ?", id).Updates(&dataUpdated)
		if updated.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message":       "Error updating data",
				"error_message": updated.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": dataUpdated,
		})
	}
}
