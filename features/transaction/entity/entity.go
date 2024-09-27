package entity

import (
	"time"

	"gorm.io/gorm"
)

type TransactionCore struct {
	Id         uint
	Nama       string
	TotalHarga int
	Parfum     []TransactionItemCore
	ParfumDetail []TransactionItemCore
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   gorm.DeletedAt
}

type TransactionItemCore struct {
	ProdukCabangId string
	Nama           string
	Foto           string
	Harga          int
	Jumlah         int
	HargaTotal     int
}
