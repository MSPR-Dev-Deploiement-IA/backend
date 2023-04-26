package handlers

import "github.com/gin-gonic/gin"

func (h Handler) HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
