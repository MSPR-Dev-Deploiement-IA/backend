package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SenderID    uint      `json:"sender_id" gorm:"not null"`
	ReceiverID  uint      `json:"receiver_id" gorm:"not null"`
	MessageText string    `json:"message_text" gorm:"type:varchar(255);not null"`
	Timestamp   time.Time `json:"timestamp" gorm:"not null"`
}

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// Create a new message
func (r *MessageRepository) Create(message *Message) error {
	return r.db.Create(message).Error
}

// Get a message by ID
func (r *MessageRepository) GetByID(id uint) (*Message, error) {
	var message Message
	err := r.db.First(&message, id).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// Update a message
func (r *MessageRepository) Update(message *Message) error {
	return r.db.Save(message).Error
}

// Delete a message
func (r *MessageRepository) Delete(message *Message) error {
	return r.db.Delete(message).Error
}

// Get all messages between two users
func (r *MessageRepository) GetByUsers(senderID uint, receiverID uint) ([]Message, error) {
	var messages []Message
	err := r.db.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", senderID, receiverID, receiverID, senderID).Order("timestamp asc").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
