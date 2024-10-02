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
func (produkRepo *produkCRepository) GetAllProduk(search, filter string) ([]entity.ProdukCabangCore, error) {
	var dataProduk []model.ProdukCabang

	query := produkRepo.db.Model(&model.ProdukCabang{})

	if search != "" {
		query = query.Where("nama_produk LIKE ?", "%"+search+"%").Order("created_at DESC")
	}

	if filter == "asc" {
		query = query.Order("nama_produk ASC")
	} else if filter == "desc" {
		query = query.Order("nama_produk DESC")
	} else {
		query = query.Order("created_at DESC")
	}

	tx := query.Order("created_at DESC").Find(&dataProduk)
	if tx.Error != nil {
		return nil, tx.Error
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

// InputRiwayat implements entity.ProdukCabangRepositoryInterface.
func (produkRepo *produkCRepository) InputRiwayat(data entity.RiwayatProdukCabangCore) (entity.RiwayatProdukCabangCore, error) {
	input := entity.RiwayatCoreToRiwayatModel(data)

	errRiwayat := produkRepo.db.Save(&input)
	if errRiwayat.Error != nil {
		return entity.RiwayatProdukCabangCore{}, errRiwayat.Error
	}

	var resp = entity.RiwayatModelToRiwayatCore(input)

	return resp, nil
}

// GetAllRiwayat implements entity.ProdukCabangRepositoryInterface.
func (produkRepo *produkCRepository) GetAllRiwayat() ([]entity.RiwayatProdukCabangCore, error) {
	var dataHistory []model.RiwayatProdukCabang

	errData := produkRepo.db.Preload("ProdukDetail").Find(&dataHistory).Error
	if errData != nil {
		return nil, errData
	}

	mapData := entity.ListHistoryModelToHistoryCore(dataHistory)

	return mapData, nil
}