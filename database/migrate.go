package database

import (
	"log"

	"github.com/harshalhonmote/crs/config"
	"github.com/harshalhonmote/crs/models"
	"gorm.io/gorm"
)

type Database struct {
}

func (db *Database) GetDB() *gorm.DB {
	return config.GetDB()
}

func Migrate() {
	log.Printf("Migrate: Start")
	db := config.GetDB()
	db.AutoMigrate(
		&models.User{},
		&models.Branch{},
		&models.Car{},
		&models.Booking{},
		&models.Transaction{},
	)
	log.Printf("Migrate: Success")
}
