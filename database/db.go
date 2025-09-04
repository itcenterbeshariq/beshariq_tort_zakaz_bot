package database

import (
	"beshariq_tort_zakaz_bot/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {

	db, err := gorm.Open(sqlite.Open("tort.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	migrate()
	seeding()
}

func migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Contact{},
	)

	if err != nil {
		log.Fatal(err)
	}
}

func seeding() {

	SeedAdmin(DB)

}
