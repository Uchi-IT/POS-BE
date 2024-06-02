package model

import (
	"time"

	"gorm.io/gorm"
)

type ProdukGudang struct {
	Id         string `gorm:"primary key"`
	Foto       string
	NamaProduk string
	Stok       int
	HargaJual  int
	HargaBeli  int
	// UserId string
	CreatedAt  time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt  time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
