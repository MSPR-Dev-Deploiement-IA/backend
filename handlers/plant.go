package handlers

import (
	"backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) AddPlant(c *gin.Context) {
	var responseJson gin.H

	fmt.Println("AddPlant")

	userId, err := c.MustGet("userID").(uint)
	if !err {
		fmt.Println("userId not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId not found"})
		return
	}

	// Example of json sent
	//{
	//	"name": "Test",
	//	"description": "aeazeaze",
	//	"newSpecies": "Tespece",
	//	"newAddress": {
	//		"address": "32 rue de cénac, C34",
	//			"zipCode": "33100",
	//			"city": "Bordeaux",
	//			"country": "FR"
	//		}
	//	}

	if err := c.ShouldBindJSON(&responseJson); err != nil {
		fmt.Println("Error while binding json")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var species models.Species
	var address models.Location

	if responseJson["newSpecies"] != nil {
		// Create new species
		species.CommonName = responseJson["newSpecies"].(string)
		err := species.GetData()
		if err != nil {
			fmt.Println("Error while getting species data")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		h.db.Create(&species)
	} else {
		// Get species
		h.db.Where("common_name = ?", responseJson["species"]).First(&species)
	}

	if responseJson["newAddress"] != nil {
		// Create new address
		address.Address = responseJson["newAddress"].(map[string]interface{})["address"].(string)
		address.ZipCode = responseJson["newAddress"].(map[string]interface{})["zipCode"].(string)
		address.City = responseJson["newAddress"].(map[string]interface{})["city"].(string)
		address.Country = responseJson["newAddress"].(map[string]interface{})["country"].(string)
		address.Name = responseJson["newAddress"].(map[string]interface{})["name"].(string)
		address.UserID = userId
		lat, lon, err := address.CalculateLatLon()
		address.Latitude = lat
		address.Longitude = lon

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// print all address fields
		fmt.Println(address)

		result := h.db.Create(&address)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	} else {
		result := h.db.Where("name = ? AND user_id = ?", responseJson["address"], userId).First(&address)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	var newPlant models.Plant
	newPlant.Name = responseJson["name"].(string)
	newPlant.Description = responseJson["description"].(string)
	newPlant.SpeciesID = species.ID
	newPlant.LocationID = address.ID
	newPlant.UserID = userId
	newPlant.Species = species
	newPlant.Location = address

	result := h.db.Create(&newPlant)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"plant_id": newPlant.ID})
}

func (h Handler) GetUserPlants(c *gin.Context) {
	var plants []models.Plant

	userId, err := c.MustGet("userID").(uint)
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId not found"})
		return
	}

	result := h.db.Where("user_id = ?", userId).Preload("Species").Preload("Location").Find(&plants)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plants": plants})
}

func (h Handler) GetPlants(c *gin.Context) {
	var plants []models.Plant

	result := h.db.Preload("Species").Preload("Location").Find(&plants)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plants": plants})
}

func (h Handler) UploadPlantFile(c *gin.Context) {
	userId := c.MustGet("userID").(uint)

	plantIdString := c.Query("plantId")
	plantId, err := strconv.Atoi(plantIdString)
	if err != nil {
		fmt.Println("Error converting plantId string to integer: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var plant models.Plant
	result := h.db.Where("id = ? AND user_id = ?", plantId, userId).First(&plant)

	if result.Error != nil {
		fmt.Println("Error retrieving plant from database: ", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if plant.ID == 0 {
		fmt.Println("No plant found with the provided id")
		c.JSON(http.StatusNotFound, gin.H{"error": "Plant not found"})
		return
	}

	//	Get multiple files and upload to /uploads in the server
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("Error retrieving multipart form: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]

	for _, file := range files {
		randomId := uuid.New().String()
		file.Filename = randomId + "-" + file.Filename
		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, "./uploads/"+file.Filename)
		if err != nil {
			fmt.Println("Error saving uploaded file: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var photo models.Photo
		photo.PhotoFileUrl = file.Filename
		photo.PlantID = &plant.ID
		photo.UserID = userId
		result := h.db.Create(&photo)

		if result.Error != nil {
			fmt.Println("Error saving photo in the database: ", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}
}

func (h Handler) GetPlantById(c *gin.Context) {
	plantIdString := c.Param("plantId")
	plantId, err := strconv.Atoi(plantIdString)
	if err != nil {
		fmt.Println("Error converting plantId string to integer: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var plant models.Plant
	result := h.db.Where("id = ?", plantId).Preload("Species").Preload("Location").Preload("Photos").First(&plant)
	if result.Error != nil {
		fmt.Println("Error retrieving plant from database: ", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if plant.ID == 0 {
		fmt.Println("No plant found with the provided id")
		c.JSON(http.StatusNotFound, gin.H{"error": "Plant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plant": plant})
}

func (h Handler) GetPlantsByUser(c *gin.Context) {
	userIdString := c.Param("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		fmt.Println("Error converting userId string to integer: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var plants []models.Plant
	// Preload PlantHistories, PlantAdvices, Photos, Location
	result := h.db.Where("user_id = ?", userId).
		Preload("Species").
		Preload("PlantHistories").
		Preload("PlantAdvices").
		Preload("Photos").
		Preload("Location").
		Find(&plants)

	if result.Error != nil {
		fmt.Println("Error retrieving plants from database: ", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plants": plants})
}
