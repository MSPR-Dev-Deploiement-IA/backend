package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) BecomeBotanist(c *gin.Context) {
	userId := c.MustGet("userID").(uint)

	var BecomeBotanist models.BecomeBotanist
	if err := c.ShouldBindJSON(&BecomeBotanist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	BecomeBotanist.UserID = userId

	if err := h.db.Create(&BecomeBotanist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, BecomeBotanist)
}
