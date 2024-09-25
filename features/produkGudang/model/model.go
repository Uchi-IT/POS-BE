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
	CreatedAt  time.Time      `gorm:"timestamp"`
	UpdatedAt  time.Time      `gorm:"timestamp"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

