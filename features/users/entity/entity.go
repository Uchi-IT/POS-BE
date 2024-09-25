package entity

import (
	"time"

	"gorm.io/gorm"
)

type UsersCore struct {
	Id        string
	Email     string
	Password  string
	Nama      string
	Cabang_id []string
	Cabang    []UserCabangCore
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt
}

type UserCabangCore struct {
	Cabang string
}
