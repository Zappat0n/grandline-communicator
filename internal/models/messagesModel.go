package models

import (
	"gorm.io/gorm"

	"albarros/grandline/communicator/internal/db"
)

type Message struct {
	gorm.Model
	Content string
	UserID  uint
}

func CreateMessage(content string, userID uint) {
	db.Grandline.Create(&Message{
		Content: content,
		UserID:  userID,
	})
}
