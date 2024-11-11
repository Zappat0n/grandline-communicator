package main

import (
	"gorm.io/gorm"

	"albarros/grandline/communicator/internal/db"
	"albarros/grandline/communicator/internal/initializers"
)

var DB *gorm.DB

func init() {
	initializers.LoadEnvVariables()
	db.ConnectGrandline()
	initializers.Routes()
}

func main() {
}
