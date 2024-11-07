package database

import (
	"log"
	"meli/models"
	"meli/seeders"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func DbConnection() *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open("database.db?_foreign_keys=on"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	errMigration := db.AutoMigrate(&models.User{}, &models.Place{}, &models.Show{}, &models.Section{}, &models.Reservation{})
	if errMigration != nil {
		log.Fatalf("Error in table migration: %v", err)
	}

	seeders.PlaceSeed(db)
	seeders.ShowSeed(db)
	seeders.SectionSeed(db)
	seeders.UserSeed(db)



	return db
}
