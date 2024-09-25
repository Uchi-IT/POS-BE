package model

import (
	"time"
	"uchiiParfume/features/produkCabang/model"

	"gorm.io/gorm"
)

type Transaction struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	Nama       string
	TotalHarga int
	Parfum     []model.ProdukCabang `gorm:"many2many:TransactionItem"`
	parfum_id  []string             `gorm:"-"`
	CreatedAt  time.Time            `gorm:"type:timestamp"`
	UpdatedAt  time.Time            `gorm:"type:timestamp"`
	DeleteAt   gorm.DeletedAt       `gorm:"index"`
}

type TransactionItem struct {
	TransactionId  uint
	ProdukCabangId string
	Nama           string
	Foto           string
	Harga          int
	Jumlah         int
	HargaTotal     int
}
