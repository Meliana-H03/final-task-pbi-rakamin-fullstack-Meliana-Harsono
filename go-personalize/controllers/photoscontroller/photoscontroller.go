package photoscontroller

import (
	"net/http"
	"time"

	//"strconv"

	"github.com/Meliana03/go-personalizeapi/models"
	"github.com/gin-gonic/gin"
)

// CreatePhoto handles creating a new photo
func CreatePhoto(c *gin.Context) {
	var photo models.Photos
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	photo.UserID = userID.(int)
	photo.CreatedAt = time.Now()
	photo.UpdatedAt = time.Now()

	if err := models.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}

// GetPhotos handles retrieving all photos
func GetPhotos(c *gin.Context) {
	var photos []models.Photos
	if err := models.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, photos)
}

// UpdatePhoto handles updating a photo
func UpdatePhoto(c *gin.Context) {
	var photo models.Photos
	if err := models.DB.First(&photo, c.Param("photoId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	userID, _ := c.Get("user_id")
	if photo.UserID != userID.(int) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to update this photo"})
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photo.UpdatedAt = time.Now()
	if err := models.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}

// DeletePhoto handles deleting a photo
func DeletePhoto(c *gin.Context) {
	var photo models.Photos
	if err := models.DB.First(&photo, c.Param("photoId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	userID, _ := c.Get("user_id")
	if photo.UserID != userID.(int) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to delete this photo"})
		return
	}

	if err := models.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
