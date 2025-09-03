package models

import "gorm.io/gorm"

const RoleAdmin = "admin"
const RoleUser = "user"

type User struct {
	FullName string `gorm:"null"`
	Role     string `gorm:"type:varchar(30);not null"`
	Password string `gorm:"type:text;not null"`

	gorm.Model
}
