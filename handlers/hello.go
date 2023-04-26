package handlers

import "github.com/gin-gonic/gin"

func (h Handler) HelloHandler(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	c.JSON(200, gin.H{
		"message": "Hello, World!",
		"user_id": userID,
	})
}
