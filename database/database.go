package database

import (
	"servertestgo/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase(direction string) {
	//para la base de datos usamos go orm, y por practicidad de este proyecto usamos el driver sql lite
	//en go orm se puede remplazar el driver por otro
	db, err := gorm.Open(sqlite.Open(direction), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.Order{})

}
