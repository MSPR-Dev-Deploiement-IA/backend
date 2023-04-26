package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) Save(db *gorm.DB) error {
	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
