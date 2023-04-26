package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (u *User) Save(db *gorm.DB) error {
	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) GetUser(db *gorm.DB) error {
	result := db.Where("email = ?", u.Email).First(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
