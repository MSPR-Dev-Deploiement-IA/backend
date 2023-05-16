package models

import (
	"gorm.io/gorm"
)

type PlantAdvice struct {
	gorm.Model
	PlantID    uint   `json:"plant_id" gorm:"not null"`
	UserID     uint   `json:"user_id" gorm:"not null"`
	AdviceText string `json:"advice_text" gorm:"type:varchar(255);not null"`
}

type PlantAdviceRepository struct {
	db *gorm.DB
}

func NewPlantAdviceRepository(db *gorm.DB) *PlantAdviceRepository {
	return &PlantAdviceRepository{db: db}
}

// Create a new plant advice
func (r *PlantAdviceRepository) Create(plantAdvice *PlantAdvice) error {
	return r.db.Create(plantAdvice).Error
}

// Get a plant advice by ID
func (r *PlantAdviceRepository) GetByID(id uint) (*PlantAdvice, error) {
	var plantAdvice PlantAdvice
	err := r.db.First(&plantAdvice, id).Error
	if err != nil {
		return nil, err
	}
	return &plantAdvice, nil
}

// Update a plant advice
func (r *PlantAdviceRepository) Update(plantAdvice *PlantAdvice) error {
	return r.db.Save(plantAdvice).Error
}

// Delete a plant advice
func (r *PlantAdviceRepository) Delete(plantAdvice *PlantAdvice) error {
	return r.db.Delete(plantAdvice).Error
}

// Get all plant advices for a plant
func (r *PlantAdviceRepository) GetByPlantID(plantID uint) ([]PlantAdvice, error) {
	var plantAdvices []PlantAdvice
	err := r.db.Where("plant_id = ?", plantID).Find(&plantAdvices).Error
	if err != nil {
		return nil, err
	}
	return plantAdvices, nil
}
