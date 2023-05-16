package models

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  uint `json:"user_id" gorm:"not null"`
	PlantID uint `json:"plant_id" gorm:"not null"`
}

type FavoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{db: db}
}

// Create a new favorite
func (r *FavoriteRepository) Create(favorite *Favorite) error {
	return r.db.Create(favorite).Error
}

// Get a favorite by ID
func (r *FavoriteRepository) GetByID(id uint) (*Favorite, error) {
	var favorite Favorite
	err := r.db.First(&favorite, id).Error
	if err != nil {
		return nil, err
	}
	return &favorite, nil
}

// Update a favorite
func (r *FavoriteRepository) Update(favorite *Favorite) error {
	return r.db.Save(favorite).Error
}

// Delete a favorite
func (r *FavoriteRepository) Delete(favorite *Favorite) error {
	return r.db.Delete(favorite).Error
}

// Get all favorites for a user
func (r *FavoriteRepository) GetByUserID(userID uint) ([]Favorite, error) {
	var favorites []Favorite
	err := r.db.Where("user_id = ?", userID).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}
