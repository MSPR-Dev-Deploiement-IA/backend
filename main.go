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
	backend := router.Group("/backend")
	backend.Static("/static", "./uploads")
	auth := backend.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.Refresh)
	}

	api := backend.Group("/api")
	api.Use(m.Authorize())
	{

		api.GET("/hello", h.HelloHandler)
		// Add more secured routes here

		users := api.Group("/users")
		{
			//users.GET("/", h.GetUsers)
			users.GET("/me", h.GetCurrentUser)
		}

		plants := api.Group("/plants")
		{
			plants.GET("/", h.GetPlants)
			plants.POST("/add", h.AddPlant)
			plants.POST("/upload", h.UploadPlantFile)
			plants.GET("/:plantId", h.GetPlantById)
			plants.GET("/user/:userId", h.GetPlantsByUser)
		}

		species := api.Group("/species")
		{
			species.GET("/", h.GetSpecies)
			species.GET("/:commonName", h.GetSpeciesByCommonName)
		}

		locations := api.Group("/locations")
		{
			locations.GET("/", h.GetUserLocations)
		}

		messages := api.Group("/messages")
		{
			messages.GET("/", h.GetAllMessages)
			messages.POST("/add", h.PostMessage)
		}

		photos := api.Group("/photos")
		{
			photos.POST("/", h.UploadFile)
			photos.GET("/", h.GetFileByUserId)
		}

		history := api.Group("/history")
		{
			history.POST("/", h.CreatePlantHistory)
			history.GET("/:id", h.GetPlantHistory)
		}
	}

	err = router.Run()
	if err != nil {
		log.Fatal(err)
	} // listen and serve on 0.0.0.0:8080 (for example)
}
