package handlers

import (
	"backend/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h Handler) CreatePlantHistory(c *gin.Context) {
	userId := c.MustGet("userID").(uint)

	var plantHistory models.PlantHistory
	err := c.BindJSON(&plantHistory)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plantHistory.UserID = userId

	err = plantHistory.Save(h.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h Handler) GetPlantHistory(c *gin.Context) {
	// Get the plant ID from the path parameters.
	plantID := c.Param("id")

	// Convert the plant ID to an integer.
	id, err := strconv.Atoi(plantID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
		return
	}

	// Create a PlantHistory object to hold the result.
	var plantHistory models.PlantHistory

	// Use GORM's First method to find the first record with the given ID.
	err = h.db.First(&plantHistory, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Plant history not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Return the plant history as JSON.
	c.JSON(http.StatusOK, plantHistory)
}
