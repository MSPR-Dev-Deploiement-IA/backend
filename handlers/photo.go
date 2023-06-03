package handlers

import (
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) UploadFile(c *gin.Context) {
	userId := c.MustGet("userID").(uint)

	//	Get multiple files and upload to /uploads in the server
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("Error retrieving multipart form: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]

	for _, file := range files {
		randomId := uuid.New().String()
		file.Filename = randomId + "-" + file.Filename
		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, "./uploads/"+file.Filename)
		if err != nil {
			fmt.Println("Error saving uploaded file: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var photo models.Photo
		photo.PhotoFileUrl = file.Filename
		photo.UserID = userId
		result := h.db.Create(&photo)

		if result.Error != nil {
			fmt.Println("Error saving photo in the database: ", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}
}

func (h Handler) GetFileByUserId(c *gin.Context) {
	userId := c.MustGet("userID").(uint)

	var photos []models.Photo
	result := h.db.Where("user_id = ?", userId).Find(&photos)
	if result.Error != nil {
		fmt.Println("Error retrieving photos: ", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photos": photos})
}
