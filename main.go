package main

import (
	"backend/database"
	"backend/handlers"
	"backend/middlewares"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file -- ignore if prod")
	}

	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	h := handlers.Newhandler(db)
	m := middlewares.NewMiddleware(db)

	router := gin.Default()

	router.Use(m.CORSMiddleware())

	router.GET("/tests", h.Tests)

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.Refresh)
	}

	api := router.Group("/api")
	api.Use(m.Authorize())
	{
		api.GET("/hello", h.HelloHandler)
		// Add more secured routes here
		api.GET("/users/me", h.GetCurrentUser)

		api.POST("/upload", h.Plant)

	}

	err = router.Run()
	if err != nil {
		log.Fatal(err)
	} // listen and serve on 0.0.0.0:8080 (for example)
}
