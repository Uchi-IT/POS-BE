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
	SeratusMl           int `gorm:"default:0"`
	DuaratusMl          int `gorm:"default:0"`
	DuaRatusLimaPuluhMl int `gorm:"default:0"`
	Manual              int `gorm:"default:0"`
	Total               int `gorm:"default:0"`
	ProdukId            string
	CabangId            string
	CreatedAt           time.Time      `gorm:"type:timestamp"`
	UpdatedAt           time.Time      `gorm:"type:timestamp"`
	DeleteAt            gorm.DeletedAt `gorm:"index"`
}
