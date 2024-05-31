package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        string `gorm:"primary key"`
	Email     string
	Password  string
	RoleId    string
	Cabang    string
	Role      string
	CreatedAt time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
