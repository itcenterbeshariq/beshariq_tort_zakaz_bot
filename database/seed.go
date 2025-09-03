package database

import (
	"beshariq_tort_zakaz_bot/models"
	"errors"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {

	var user models.User
	if err := db.Where("role = ?", models.RoleAdmin).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.FullName = "Admin"
			user.Role = models.RoleAdmin
			user.Password = models.RoleAdmin
			db.Create(&user)
		}
	}
}
