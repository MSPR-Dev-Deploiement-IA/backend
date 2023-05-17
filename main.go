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

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.Refresh)
	}

	api := router.Group("/api")
	api.Use(h.Authorize())
	{
		api.GET("/hello", h.HelloHandler)
		// Add more secured routes here
	}

	err = router.Run()
	if err != nil {
		log.Fatal(err)
	} // listen and serve on 0.0.0.0:8080 (for example)
}
