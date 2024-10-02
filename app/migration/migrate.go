package migration

import (
	users "uchiiParfume/features/users/model"
	cabang "uchiiParfume/features/cabang/model"
	produkCabang "uchiiParfume/features/produkCabang/model"
	produkGudang "uchiiParfume/features/produkGudang/model"
	transaction "uchiiParfume/features/transaction/model"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB){
	db.AutoMigrate(&cabang.Cabang{})
	db.AutoMigrate(&users.User{}, &users.UserCabang{})
	db.AutoMigrate(&produkCabang.ProdukCabang{}, &produkCabang.RiwayatProdukCabang{}, &produkCabang.RiwayatProdukCabangItem{})
	db.AutoMigrate(&produkGudang.ProdukGudang{})
	db.AutoMigrate(&transaction.Transaction{}, &transaction.TransactionItem{})
}