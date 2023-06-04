package models

import "gorm.io/gorm"

type BecomeBotanist struct {
	gorm.Model
	UserID uint   `json:"userID"`
	Text   string `json:"text"`
}
