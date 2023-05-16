package models

import (
	"gorm.io/gorm"
)

type PlantAdvice struct {
	gorm.Model
	PlantID    uint   `gorm:"not null"`
	UserID     uint   `gorm:"not null"`
	AdviceText string `gorm:"type:varchar(255);not null"`
}
