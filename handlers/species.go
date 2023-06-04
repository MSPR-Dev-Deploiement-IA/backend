package handlers

import (
	"backend/models"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetSpecies(c *gin.Context) {
	var species []models.Species
	h.db.Find(&species)

	c.JSON(200, gin.H{"species": species})
}

func (h Handler) GetSpeciesByCommonName(c *gin.Context) {
	var species models.Species
	commonName := c.Param("commonName")
	h.db.Where("common_name = ?", commonName).First(&species)

	if species.ID == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Species not found"})
		return
	}

	c.JSON(200, gin.H{"species": species})
}

func (h Handler) GetSpeciesByScientific(c *gin.Context) {
	var species models.Species
	scientific := c.Param("scientific")
	h.db.Where("scientific = ?", scientific).Preload("SpeciesAdvice").First(&species)

	if species.ID == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Species not found"})
		return
	}

	c.JSON(200, gin.H{"species": species})
}
