package entity

import (
	"time"

	"gorm.io/gorm"
)

type UsersCore struct {
	Id        string
	Email     string
	Password  string
	RoleId    string
	Cabang    string
	Role      string
	// Cabang_id []string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt
}
