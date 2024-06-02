package model

import (
	"time"

	"gorm.io/gorm"
)

type ProdukCabang struct {
	Id                  string `gorm:"primary key"`
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
	CreatedAt time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time      `gorm:"type:DATETIME(0)"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}
