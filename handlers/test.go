package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) Tests(c *gin.Context) {
	var location models.Location
	err := c.ShouldBindJSON(&location)
	if err != nil {
		c.JSON(400, gin.H{"error binding JSON": err.Error()})
		return
	}

	err = location.CalculateLatLon()
	if err != nil {
		c.JSON(400, gin.H{"error calculating lat lon": err.Error()})
		return
	}

	tx := h.db.Create(&location)
	if tx.Error != nil {
		c.JSON(400, gin.H{"error creating location": tx.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"location": location})
}
