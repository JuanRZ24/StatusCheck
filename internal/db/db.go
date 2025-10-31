package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)


func Connect () *gorm.DB {
	db, err := gorm.Open(sqlite.Open("status.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos", err)
	}
	log.Println("Base de datos conectada")
	return db
}
