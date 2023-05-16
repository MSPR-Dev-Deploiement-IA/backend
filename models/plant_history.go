package models

import (
	"time"

	"gorm.io/gorm"
)

type PlantHistory struct {
	gorm.Model
	PlantID   uint      `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
}

func (p *PlantHistory) Save(db *gorm.DB) error {
	result := db.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
