package handlers

import "gorm.io/gorm"

type Handler struct {
	db *gorm.DB
}

func Newhandler(db *gorm.DB) Handler {
	return Handler{
		db: db,
	}
}
