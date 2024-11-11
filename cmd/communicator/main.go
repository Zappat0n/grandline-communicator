package main

import (
	"gorm.io/gorm"

	"albarros/grandline/communicator/internal/db"
	"albarros/grandline/communicator/internal/initializers"
)

var DB *gorm.DB

func init() {
	db.ConnectGrandline()
	initializers.Routes()
}

func main() {
}
