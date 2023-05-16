package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fetch current user's information
func (h Handler) GetUser(c *gin.Context) {
	// Retrieve user id from context
	userID := c.MustGet("userID").(uint)

	var user models.User
	db := h.db.First(&user, userID)
	if db.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Update current user's information
func (h Handler) UpdateUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = userID
	if err := user.Save(h.db); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Fetch plants a user has cared for
func (h Handler) GetCaredPlants(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var histories []models.PlantHistory
	db := h.db.Where("caretaker_id = ?", userID).Find(&histories)
	if db.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plants": histories})
}

// Fetch plants that have been cared for by others
func (h Handler) GetPlantsCaredByOthers(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var histories []models.PlantHistory
	db := h.db.Where("owner_id = ?", userID).Find(&histories)
	if db.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plants": histories})
}

// Fetch plants to be cared for
func (h Handler) GetPlantsToBeCared(c *gin.Context) {
	var plants []models.Plant
	db := h.db.Where("owner_id IS NULL").Find(&plants)
	if db.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}
}
