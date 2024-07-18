package entity

import (
	"time"

	"gorm.io/gorm"
)

type CabangCore struct {
	Id         string
	Image      string
	NamaCabang string
	Alamat     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   gorm.DeletedAt
}
