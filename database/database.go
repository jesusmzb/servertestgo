package database

import (
	"servertestgo/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDatabase(direction string) {
	//para la base de datos usamos go orm, y por practicidad de este proyecto usamos el driver sql lite
	//en go orm se puede remplazar el driver por otro
	db, err = gorm.Open(sqlite.Open(direction), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.OrderBookEntity{})
	db.AutoMigrate(&models.OrderEntity{})

}

func OrderBookStore(model models.OrderBookEntity) {
	db.Create(&model)
}

func OrderStore(model models.OrderEntity) {
	db.Create(&model)
}
