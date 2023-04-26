package main

import (
	"backend/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_, err = database.SetupDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for example)
}
