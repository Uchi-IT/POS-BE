package model

import "time"

type Notifikasi struct {
	Id string `gorm:"primary key"`
	TotalStock string
	Message string
	// produkGudangId string
	// CabangId   string
	CreatedAt  time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt  time.Time `gorm:"type:DATETIME(0)"`
}
