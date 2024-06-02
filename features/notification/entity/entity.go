package entity

import "time"

type NotifikasiCore struct {
	Id string
	// produkGudangId string
	TotalStock string
	CabangId   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
