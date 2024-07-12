package migration

import (
	users "uchiiParfume/features/users/model"
	

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB){
	db.AutoMigrate(&users.User{})
}