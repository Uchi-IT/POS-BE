package repository

import (
	"errors"
	"uchiiParfume/features/produkCabang/entity"
	"uchiiParfume/features/produkCabang/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type produkCRepository struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) entity.ProdukCabangRepositoryInterface {
	return &produkCRepository{
		db: db,
	}
}

// DeleteProduk implements entity.ProdukCabangRepositoryInterface.
func (produkRepo *produkCRepository) DeleteProduk(id string) error {
	var chekcId model.ProdukCabang

	errData := produkRepo.db.Where("id = ?", id).Delete(&chekcId)
	if errData != nil {
		return errData.Error
	}

	return nil
}

// GetAllProduk implements entity.ProdukCabangRepositoryInterface.
func (produkRepo *produkCRepository) GetAllProduk() ([]entity.ProdukCabangCore, error) {
	var dataProduk []model.ProdukCabang

	errData := produkRepo.db.Find(&dataProduk).Error
	if errData != nil {
		return nil, errData
	}

	mapData := entity.ListProdukCabangModelToListProdukCabangCore(dataProduk)
	return mapData, nil
}

// GetById implements entity.ProdukCabangRepositoryInterface.
func (produkRepo *produkCRepository) GetById(id string) (entity.ProdukCabangCore, error) {
	var data model.ProdukCabang
	errData := produkRepo.db.Where("id=?", id).First(&data).Error
	if errData != nil {
		return entity.ProdukCabangCore{}, errData
	}

	resp := entity.ProdukCabangModelToProdukCabangCore(data)

	return resp, nil
}

// InputProduk implements entity.ProdukCabangRepositoryInterface.
func (produkRepo *produkCRepository) InputProduk(data entity.ProdukCabangCore) (entity.ProdukCabangCore, error) {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return entity.ProdukCabangCore{}, UUIDerr
	}

	input := entity.ProdukCabangCoreToProdukCabangModel(data)
	input.Id = newUUID.String()

	errCabang := produkRepo.db.Save(&input)
	if errCabang.Error != nil {
		return entity.ProdukCabangCore{}, errCabang.Error
	}

	var resp = entity.ProdukCabangModelToProdukCabangCore(input)

	return resp, nil
}

// UpdateProduk implements entity.ProdukCabangRepositoryInterface.
func (produkRepo *produkCRepository) UpdateProduk(id string, data entity.ProdukCabangCore) error {
	request := entity.ProdukCabangCoreToProdukCabangModel(data)

	tx := produkRepo.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}
