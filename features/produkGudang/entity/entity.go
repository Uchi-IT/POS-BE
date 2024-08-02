package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProdukGudangCore struct {
	Id         string
	Foto       string
	NamaProduk string
	Stok       int
	HargaJual  int
	HargaBeli  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   gorm.DeletedAt
}
