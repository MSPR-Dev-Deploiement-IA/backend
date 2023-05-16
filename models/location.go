package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	UserID    uint    `gorm:"not null"`
	Latitude  float64 `gorm:"not null"`
	Longitude float64 `gorm:"not null"`
}
