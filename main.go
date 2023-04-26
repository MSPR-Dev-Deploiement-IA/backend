package main

import (
	"backend/database"
	"backend/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	router := gin.Default()

	h := handlers.Newhandler(db)

	router.GET("/", h.HelloHandler)
	router.POST("/users", h.PostUser)

	router.Run() // listen and serve on 0.0.0.0:8080 (for example)

}
