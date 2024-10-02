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
	Total               int
	ProdukId            string
	CabangId            string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeleteAt            gorm.DeletedAt
}

type RiwayatProdukCabangCore struct {
	Id           uint 
	NamaCabang   string
	NamaAdmin    string
	ProdukDetail []RiwayatProdukCabangItemCore
	TotalStok    int
	Catatan      string
	CreatedAt    time.Time      `gorm:"type:timestamp"`
	UpdatedAt    time.Time      `gorm:"type:timestamp"`
	DeleteAt     gorm.DeletedAt `gorm:"index"`
}

type RiwayatProdukCabangItemCore struct {
	Id                  string
	RiwayatId           uint
	ProdukCabangId      string
	NamaProduk          string
	SeratusMl           int `gorm:"default:0"`
	DuaratusMl          int `gorm:"default:0"`
	DuaRatusLimaPuluhMl int `gorm:"default:0"`
	Manual              int `gorm:"default:0"`
}
