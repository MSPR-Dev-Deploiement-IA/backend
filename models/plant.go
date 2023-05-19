package models

import (
	"gorm.io/gorm"
)

type Plant struct {
	gorm.Model
	SpeciesID      uint           `json:"species_id"`
	Species        Species        `json:"species" gorm:"foreignKey:SpeciesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlantHistories []PlantHistory `json:"plant_histories" gorm:"foreignKey:PlantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlantAdvices   []Advice       `json:"plant_advices" gorm:"foreignKey:PlantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Photos         []Photo        `json:"photos" gorm:"foreignKey:PlantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LocationID     uint           `json:"location_id"`
	Location       Location       `json:"location" gorm:"foreignKey:LocationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID         uint           `json:"user_id"` // Add this line
	Description    string         `json:"description"`
}

func (p *Plant) Save(db *gorm.DB) error {
	result := db.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type PlantRepository struct {
	db *gorm.DB
}

func NewPlantRepository(db *gorm.DB) *PlantRepository {
	return &PlantRepository{db: db}
}

// Create a new plant
func (r *PlantRepository) Create(plant *Plant) error {
	return r.db.Create(plant).Error
}

// Get a plant by ID
func (r *PlantRepository) GetByID(id uint) (*Plant, error) {
	var plant Plant
	err := r.db.Preload("PlantHistories").Preload("PlantAdvices").Preload("Photos").First(&plant, id).Error
	if err != nil {
		return nil, err
	}
	return &plant, nil
}

// Update a plant
func (r *PlantRepository) Update(plant *Plant) error {
	return r.db.Save(plant).Error
}

// Delete a plant
func (r *PlantRepository) Delete(plant *Plant) error {
	return r.db.Delete(plant).Error
}

// Get all plants for a user
func (r *PlantRepository) GetByUserID(userID uint) ([]Plant, error) {
	var plants []Plant
	err := r.db.Where("user_id = ?", userID).Find(&plants).Error
	if err != nil {
		return nil, err
	}
	return plants, nil
}
