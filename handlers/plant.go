package handlers

import (
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Plant(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Save the file to a specific destination
	// Here, the file is being saved in the current directory
	if err := c.SaveUploadedFile(file, file.Filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func (h Handler) AddPlant(c *gin.Context) {
	var responseJson gin.H

	userId, err := c.MustGet("userID").(uint)
	if !err {
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
		err := address.CalculateLatLon()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		h.db.Create(&address)
	} else {
		h.db.Where("address = ?", responseJson["address"]).First(&address)
	}

	var newPlant models.Plant
	newPlant.Name = responseJson["name"].(string)
	newPlant.Description = responseJson["description"].(string)
	newPlant.Species = species
	newPlant.Location = address
	h.db.Create(&newPlant)

	c.JSON(http.StatusCreated, gin.H{"message": "File uploaded successfully"})
}
