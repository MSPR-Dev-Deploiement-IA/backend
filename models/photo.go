package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	PlantID      uint   `json:"plant_id"`
	Plant        Plant  `json:"plant" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID       uint   `json:"user_id"`
	User         User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PhotoFileUrl string `json:"photo_file_url" gorm:"type:varchar(255);not null,unique"`
}

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{db: db}
}

// Create a new photo
func (r *PhotoRepository) Create(photo *Photo) error {
	return r.db.Create(photo).Error
}

// Get a photo by ID
func (r *PhotoRepository) GetByID(id uint) (*Photo, error) {
	var photo Photo
	err := r.db.First(&photo, id).Error
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

// Update a photo
func (r *PhotoRepository) Update(photo *Photo) error {
	return r.db.Save(photo).Error
}

// Delete a photo
func (r *PhotoRepository) Delete(photo *Photo) error {
	return r.db.Delete(photo).Error
}

// Get all photos for a plant
func (r *PhotoRepository) GetByPlantID(plantID uint) ([]Photo, error) {
	var photos []Photo
	err := r.db.Where("plant_id = ?", plantID).Find(&photos).Error
	if err != nil {
		return nil, err
	}
	return photos, nil
}
