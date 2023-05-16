package models

import (
	"gorm.io/gorm"
)

type Plant struct {
	gorm.Model
	Name             string `gorm:"type:varchar(100);not null"`
	Type             string `gorm:"type:varchar(100);not null"`
	Description      string `gorm:"type:varchar(255);not null"`
	CareInstructions string `gorm:"type:varchar(255);not null"`
	UserID           uint   `gorm:"not null"`
	PlantHistories   []PlantHistory
	PlantAdvices     []PlantAdvice
	Photos           []Photo
}

func (p *Plant) Save(db *gorm.DB) error {
	result := db.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
