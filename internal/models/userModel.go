package models

import (
	"fmt"

	"gorm.io/gorm"

	"albarros/grandline/communicator/internal/db"
)

type User struct {
	gorm.Model
	Name  string
	Phone string
}

func FindUserByPhone(phone string) User {
	var user User
	result := db.Grandline.Where("phone = ?", phone).First(&user)

	if result.Error != nil {
		fmt.Println("Phone not found, ?", phone)
	}

	return user
}
