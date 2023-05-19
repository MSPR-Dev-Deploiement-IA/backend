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

	lat, lon, err := models.CalculateAndSaveLatLon(location)
	if err != nil {
		c.JSON(400, gin.H{"error calculating lat lon": err.Error()})
		return
	}

	location.Latitude = lat
	location.Longitude = lon

	tx := h.db.Create(&location)
	if tx.Error != nil {
		c.JSON(400, gin.H{"error creating location": tx.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"location": location})
}
