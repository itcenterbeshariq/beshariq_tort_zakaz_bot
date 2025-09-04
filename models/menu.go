package models

import "time"

type Menu struct {
	ID        int64
	UserID    int64
	Amount    int
	CreatedAt time.Time
}
