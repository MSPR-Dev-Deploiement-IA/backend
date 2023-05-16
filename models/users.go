package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `gorm:"type:varchar(100);not null;json:"name"`
	Email         string `gorm:"type:varchar(100);unique;not null;json:"email"`
	Password      string `gorm:"not null;json:"password"`
	RememberToken string
	Location      Location
	Plants        []Plant
	Favorites     []Favorite
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

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create a new user
func (r *UserRepository) Create(user *User) error {
	return user.Save(r.db)
}

// Get a user by ID
func (r *UserRepository) GetByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update a user
func (r *UserRepository) Update(user *User) error {
	return user.Save(r.db)
}

// Delete a user
func (r *UserRepository) Delete(user *User) error {
	return r.db.Delete(user).Error
}

// Get a user by email
func (r *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	user.Email = email
	err := user.GetUser(r.db)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
