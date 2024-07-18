package migration

import (
	users "uchiiParfume/features/users/model"
	cabang "uchiiParfume/features/cabang/model"
	produkCabang "uchiiParfume/features/produkCabang/model"
	

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB){
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&cabang.Cabang{})
	db.AutoMigrate(&produkCabang.ProdukCabang{})
}