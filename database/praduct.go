package database

import (
	"beshariq_tort_zakaz_bot/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/phonebook?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ DB ulanmadi:", err)
	}

	// Jadval yaratish (agar bo‘lmasa)
	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatal("❌ Jadval yaratilmadi:", err)
	}

	log.Println("✅ Database ulandi va migratsiya qilindi")
}
