package models

import "time"

type Product struct {
	ImageURL    string    // Surat linki
	ID          int64     // Mahsulot ID
	Weight      float64   // Og‘irligi (kg yoki litr ko‘rinishida)
	Price       int       // Narxi (so‘mda)
	Description string    // Mahsulot haqida izoh
	CreatedAt   time.Time // Yaratilgan vaqt

}
