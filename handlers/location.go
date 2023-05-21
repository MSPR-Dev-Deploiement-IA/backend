package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetUserLocations(c *gin.Context) {
	userID, _ := c.Get("userID")

	var locations []models.Location
	h.db.Where("user_id = ?", userID).Find(&locations)

	c.JSON(200, gin.H{"locations": locations})
}
