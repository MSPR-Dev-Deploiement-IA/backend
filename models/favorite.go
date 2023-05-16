package models

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  uint `gorm:"not null"`
	PlantID uint `gorm:"not null"`
}
