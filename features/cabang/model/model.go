package model

import (
	"time"

	"gorm.io/gorm"
)

type Cabang struct {
	Id         string `gorm:"primary key"`
	Image      string
	NamaCabang string
	Alamat     string
	CreatedAt  time.Time      `gorm:"type:timestamp"`
	UpdatedAt  time.Time      `gorm:"type:timestamp"`
	DeleteAt   gorm.DeletedAt `gorm:"index"`
}
