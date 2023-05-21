package handlers

import (
	"backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetAllMessages(c *gin.Context) {
	var messages []models.Message
	h.db.Preload("User").Find(&messages)

	c.JSON(200, gin.H{"messages": messages})
}

func (h Handler) PostMessage(c *gin.Context) {
	var message models.Message

	userId, err := c.MustGet("userID").(uint)

	if !err {
		fmt.Println("No user id")
		c.JSON(400, gin.H{"error": "userId not found"})
		return
	}

	if err := c.ShouldBindJSON(&message); err != nil {
		fmt.Println("Couldnt parse json")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	message.UserID = userId

	fmt.Println(message)

	if err := h.db.Create(&message).Error; err != nil {
		fmt.Println("Couldnt insert")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": message})
}
