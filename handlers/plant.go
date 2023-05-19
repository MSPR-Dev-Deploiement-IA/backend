package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Plant(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Save the file to a specific destination
	// Here, the file is being saved in the current directory
	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
