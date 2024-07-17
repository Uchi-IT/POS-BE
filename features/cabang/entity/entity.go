package entity

import "time"

type CabangCore struct {
	Id         string
	Image      string
	NamaCabang string
	Alamat     string
	UpdatedAt  time.Time
}
