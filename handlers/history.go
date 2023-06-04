package handlers

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	// Create a slice of PlantHistory objects to hold the result.
	var plantHistories []models.PlantHistory

	// Use GORM's Find method to find all records with the given plant ID.
	err = h.db.Find(&plantHistories, "plant_id = ?", id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If no records were found, return a 404.
	if len(plantHistories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No plant history found"})
		return
	}

	// Return the plant histories as JSON.
	c.JSON(http.StatusOK, plantHistories)
}
