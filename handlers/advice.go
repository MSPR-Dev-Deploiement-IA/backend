package handlers

import (
	"backend/models"

	"github.com/gin-gonic/gin"
)

func (h Handler) PostAdvice(c *gin.Context) {
	userId := c.MustGet("userID").(uint)

	var advice models.Advice
	if err := c.ShouldBindJSON(&advice); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	advice.UserID = userId

	if err := h.db.Create(&advice).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, advice)
}
