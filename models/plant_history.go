package models

import (
	"time"

	"gorm.io/gorm"
)

type PlantHistory struct {
	gorm.Model
	PlantID   uint      `json:"plant_id"`
	UserID    uint      `json:"user_id"`
	StartDate time.Time `json:"start_date" gorm:"not null"`
	EndDate   time.Time `json:"end_date" gorm:"not null"`
}

func (p *PlantHistory) Save(db *gorm.DB) error {
	result := db.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type PlantHistoryRepository struct {
	db *gorm.DB
}

func NewPlantHistoryRepository(db *gorm.DB) *PlantHistoryRepository {
	return &PlantHistoryRepository{db: db}
}

// Create a new plant history
func (r *PlantHistoryRepository) Create(plantHistory *PlantHistory) error {
	return r.db.Create(plantHistory).Error
}

// Get a plant history by ID
func (r *PlantHistoryRepository) GetByID(id uint) (*PlantHistory, error) {
	var plantHistory PlantHistory
	err := r.db.First(&plantHistory, id).Error
	if err != nil {
		return nil, err
	}
	return &plantHistory, nil
}

// Update a plant history
func (r *PlantHistoryRepository) Update(plantHistory *PlantHistory) error {
	return r.db.Save(plantHistory).Error
}

// Delete a plant history
func (r *PlantHistoryRepository) Delete(plantHistory *PlantHistory) error {
	return r.db.Delete(plantHistory).Error
}

// Get all plant histories for a plant
func (r *PlantHistoryRepository) GetByPlantID(plantID uint) ([]PlantHistory, error) {
	var plantHistories []PlantHistory
	err := r.db.Where("plant_id = ?", plantID).Find(&plantHistories).Error
	if err != nil {
		return nil, err
	}
	return plantHistories, nil
}
