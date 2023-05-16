package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SenderID    uint      `gorm:"not null"`
	ReceiverID  uint      `gorm:"not null"`
	MessageText string    `gorm:"type:varchar(255);not null"`
	Timestamp   time.Time `gorm:"not null"`
}
