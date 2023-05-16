package models

import (
	"gorm.io/gorm"
)

type Plant struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     uint   `json:"owner_id"`
}

func (p *Plant) Save(db *gorm.DB) error {
	result := db.Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
