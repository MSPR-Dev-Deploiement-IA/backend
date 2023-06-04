package models

type Advice struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Text      string `json:"text" gorm:"not null"`
	SpeciesID uint   `json:"species_id"`
	PlantID   *uint  `json:"plant_id" gorm:"type:bigint"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
