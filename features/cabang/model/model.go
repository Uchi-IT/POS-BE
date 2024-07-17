package model

import "time"

type Cabang struct {
	Id         string `gorm:"primary key"`
	Image      string
	NamaCabang string
	Alamat     string
	UpdatedAt  time.Time `gorm:"type:timestamp"`
}
