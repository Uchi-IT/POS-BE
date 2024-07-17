package model

import (
	"time"
	"uchiiParfume/features/cabang/model"

	"gorm.io/gorm"
)

type User struct {
	Id        string `gorm:"primary key"`
	Email     string
	Password  string
	Cabang    []model.Cabang `gorm:"many2many:UserCabang"`
	cabang_id []string       `gorm:"-"`
	Role      string         `gorm:"default:user"`
	CreatedAt time.Time      `gorm:"type:timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserCabang struct {
	CabangID string
	UserID   string
}
