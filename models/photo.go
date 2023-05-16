package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	PlantID   uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	PhotoFile string `gorm:"type:varchar(255);not null"`
}
