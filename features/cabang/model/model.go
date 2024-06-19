package model

import "time"

type Cabang struct {
	Id         string `gorm:"primary key"`
	NamaCabang string
	Alamat     string
	UpdatedAt  time.Time `gorm:"type:DATETIME(0)"`
}
