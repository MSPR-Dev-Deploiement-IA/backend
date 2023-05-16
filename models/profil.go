package models

import (
	"time"

	"gorm.io/gorm"
)

type Plant struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     uint   `json:"owner_id"`
}

type PlantHistory struct {
	gorm.Model
	PlantID   uint      `json:"plant_id"`
	Caretaker User      `json:"caretaker"`
	Owner     User      `json:"owner"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (p *Plant) Save(db *gorm.DB) error {
	result := db.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PlantHistory) Save(db *gorm.DB) error {
	result := db.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
