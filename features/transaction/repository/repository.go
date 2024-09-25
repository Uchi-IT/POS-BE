package repository

import (
	"errors"
	parfum "uchiiParfume/features/produkCabang/entity"
	"uchiiParfume/features/transaction/entity"
	"uchiiParfume/features/transaction/model"
	user "uchiiParfume/features/users/entity"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db     *gorm.DB
	parfum parfum.ProdukCabangRepositoryInterface
	user   user.UsersRepositoryInterface
}

func NewTransactionRepository(db *gorm.DB, parfum parfum.ProdukCabangRepositoryInterface, user user.UsersRepositoryInterface) entity.TransactionRepositoryInterface {
	return &transactionRepository{
		db:     db,
		parfum: parfum,
		user:   user,
	}
}

// CreateTransaction implements entity.TransactionRepositoryInterface.
func (transRepo *transactionRepository) CreateTransaction(data entity.TransactionCore, userId string) (entity.TransactionCore, error) {
	userData, _ := transRepo.user.GetById(userId)

	var countTotal int

	data.Nama = userData.Nama

	transactionData := entity.TransactionCoreToTransactionModel(data)

	txOuter := transRepo.db.Begin()

	if err := txOuter.Save(&transactionData).Error; err != nil {
		txOuter.Rollback()
		return entity.TransactionCore{}, err
	}

	input := entity.TransactionModelToTransactionCore(transactionData)

	for _, parfumItem := range data.Parfum {
		itemData, tx := transRepo.parfum.GetById(parfumItem.ProdukCabangId)
		if tx != nil {
			txOuter.Rollback()
			return entity.TransactionCore{}, errors.New("item tidak ada")
		}

		dataItem := new(model.TransactionItem)
		dataItem.ProdukCabangId = parfumItem.ProdukCabangId
		dataItem.TransactionId = transactionData.Id
		dataItem.Foto = itemData.Foto
		dataItem.Harga = itemData.HargaJual
		dataItem.Nama = itemData.NamaProduk
		dataItem.Jumlah = parfumItem.Jumlah
		dataItem.HargaTotal = itemData.HargaJual * parfumItem.Jumlah

		countTotal += dataItem.HargaTotal

		txInner := txOuter.Create(&dataItem)
		if txInner.Error != nil {
			txOuter.Rollback()
			return entity.TransactionCore{}, txInner.Error
		}
	}

	transactionData.TotalHarga = countTotal
	if err := txOuter.Save(&transactionData).Error; err != nil {
		txOuter.Rollback()
		return entity.TransactionCore{}, err
	}

	if err := txOuter.Exec("DELETE FROM transaction_items WHERE nama IS NULL").Error; err != nil {
		txOuter.Rollback()
		return entity.TransactionCore{}, err
	}

	data.TotalHarga = countTotal

	txOuter.Commit()

	var resp = entity.TransactionCore{
		Id:     input.Id,
		Nama:   input.Nama,
		Parfum: input.Parfum,
	}

	return resp, nil
}

// GetAllTransaction implements entity.TransactionRepositoryInterface.
func (transRepo *transactionRepository) GetAllTransaction() ([]entity.TransactionCore, error) {
	var dataTransaction []model.Transaction

	errData := transRepo.db.Preload("Parfum").Find(&dataTransaction).Error
	if errData != nil {
		return nil, errData
	}

	mapData := entity.ListTransactionModelToTransactionCore(dataTransaction)
	return mapData, nil
}

// GetAllTransactionItemById implements entity.TransactionRepositoryInterface.
func (transRepo *transactionRepository) GetAllTransactionItemById(id int) ([]entity.TransactionItemCore, error) {
	var dataTransactionItem []model.TransactionItem

	errData := transRepo.db.Where("transaction_id=?", id).Find(&dataTransactionItem).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]entity.TransactionItemCore, len(dataTransactionItem))
	for i, v := range dataTransactionItem {
		mapData[i] = entity.TransactionItemCore{
			ProdukCabangId: v.ProdukCabangId,
			Nama:           v.Nama,
			Foto:           v.Foto,
			Harga:          v.Harga,
			Jumlah:         v.Jumlah,
			HargaTotal:     v.HargaTotal,
		}
	}

	return mapData, nil
}

// GetSpecificTransaction implements entity.TransactionRepositoryInterface.
func (transRepo *transactionRepository) GetSpecificTransaction(id int) (entity.TransactionCore, error) {
	var data model.Transaction
	errData := transRepo.db.Preload("Parfum").Where("id=?", id).First(&data).Error
	if errData != nil {
		return entity.TransactionCore{}, errData
	}

	parfumItems := make([]entity.TransactionItemCore, len(data.Parfum))
	for i, v := range data.Parfum {
		parfumItems[i] = entity.TransactionItemCore{
			ProdukCabangId: v.Id,
			Nama:           v.NamaProduk,
			Foto:           v.Foto,
			Harga:          v.HargaJual,
		}
	}

	resp := entity.TransactionCore{
		Id:         data.Id,
		Nama:       data.Nama,
		TotalHarga: data.TotalHarga,
		Parfum:     parfumItems,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeleteAt:   data.DeleteAt,
	}

	return resp, nil
}