package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Grandline *gorm.DB

func ConnectGrandline() {
	var err error
	dsn := os.Getenv("DSN_URL")
	Grandline, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
