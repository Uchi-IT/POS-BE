package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProdukCabangCore struct {
	Id                  string
	Foto                string
	NamaProduk          string
	HargaJual           int
	SeratusMl           int
	DuaratusMl          int
	DuaRatusLimaPuluhMl int
	Manual              int
	// ProdukId string
	// UserId string
	// CabangId string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt
}
