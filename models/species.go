package models

type Species struct {
	ID            uint     `json:"id" gorm:"primaryKey"`
	CommonName    string   `json:"common_name" gorm:"not null"`
	Scientific    string   `json:"scientific" gorm:"not null"`
	Description   string   `json:"description" gorm:"not null"`
	SpeciesAdvice []Advice `json:"species_advice" gorm:"foreignKey:SpeciesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
