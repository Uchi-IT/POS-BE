package migration

import (
	users "uchiiParfume/features/users/model"
	cabang "uchiiParfume/features/cabang/model"
	

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB){
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&cabang.Cabang{})
}