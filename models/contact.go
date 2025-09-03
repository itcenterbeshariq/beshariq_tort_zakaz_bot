package models

import "time"

type Contact struct {
	ID            int64     `gorm:"primaryKey"` // asosiy kalit
	OwnerLastName string    `gorm:"column:owner_last_name"`
	PhoneNumber   string    `gorm:"column:phone_number"`
	Address       string    `gorm:"column:address"`
	WorkingHours  string    `gorm:"column:working_hours"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}
