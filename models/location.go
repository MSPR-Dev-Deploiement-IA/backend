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

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

// Create a new location
func (r *LocationRepository) Create(location *Location) error {
	return r.db.Create(location).Error
}

// Get a location by ID
func (r *LocationRepository) GetByID(id uint) (*Location, error) {
	var location Location
	err := r.db.First(&location, id).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

// Update a location
func (r *LocationRepository) Update(location *Location) error {
	return r.db.Save(location).Error
}

// Delete a location
func (r *LocationRepository) Delete(location *Location) error {
	return r.db.Delete(location).Error
}

// Get location by user ID
func (r *LocationRepository) GetByUserID(userID uint) (*Location, error) {
	var location Location
	err := r.db.Where("user_id = ?", userID).First(&location).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}
