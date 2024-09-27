package repository

import (
	"errors"
	"uchiiParfume/features/produkGudang/entity"
	"uchiiParfume/features/produkGudang/model"
	"uchiiParfume/utils/pagination"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type produkGRepository struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) entity.ProdukGudangRepositoryInterface {
	return &produkGRepository{
		db: db,
	}
}

// DeleteProduk implements entity.ProdukGudangRepositoryInterface.
func (produkRepo *produkGRepository) DeleteProduk(id string) error {
	var chekcId model.ProdukGudang

	errData := produkRepo.db.Where("id = ?", id).Delete(&chekcId)
	if errData != nil {
		return errData.Error
	}

	return nil
}

// GetAllProduk implements entity.ProdukGudangRepositoryInterface.
func (produkRepo *produkGRepository) GetAllProduk(page, limit int, search, filter string) ([]entity.ProdukGudangCore, pagination.PageInfo, int, error) {
	var dataProduk []model.ProdukGudang
	var totalCount int64
	offset := (page - 1) * limit

	// Initialize the query
	query := produkRepo.db.Model(&model.ProdukGudang{})

	// Add search functionality
	if search != "" {
		query = query.Where("nama_produk ILIKE ?", "%"+search+"%")
	}

	// Add filter if necessary (filter logic can be added here if needed)
	// Example: query = query.Where("some_column = ?", filter)

	// Count the total number of items before applying pagination
	tx := query.Count(&totalCount)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	// Apply offset and limit for pagination
	tx = query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&dataProduk)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	// Map the data from model to core
	mapData := entity.ListProdukGModelToProdukGCore(dataProduk)

	// Create pagination info
	pageInfo := pagination.CalculateData(int(totalCount), limit, page)

	return mapData, pageInfo, int(totalCount), nil
}


// GetById implements entity.ProdukGudangRepositoryInterface.
func (produkRepo *produkGRepository) GetById(id string) (entity.ProdukGudangCore, error) {
	var data model.ProdukGudang
	errData := produkRepo.db.Where("id=?", id).First(&data).Error
	if errData != nil {
		return entity.ProdukGudangCore{}, errData
	}

	resp := entity.ProdukGModelToProdukGCore(data)

	return resp, nil
}

// InputProduk implements entity.ProdukGudangRepositoryInterface.
func (produkRepo *produkGRepository) InputProduk(data entity.ProdukGudangCore) (entity.ProdukGudangCore, error) {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return entity.ProdukGudangCore{}, UUIDerr
	}

	input := entity.ProdukGCoreToProdukGModel(data)
	input.Id = newUUID.String()

	errCabang := produkRepo.db.Save(&input)
	if errCabang.Error != nil {
		return entity.ProdukGudangCore{}, errCabang.Error
	}

	var resp = entity.ProdukGModelToProdukGCore(input)

	return resp, nil
}

// UpdateProduk implements entity.ProdukGudangRepositoryInterface.
func (produkRepo *produkGRepository) UpdateProduk(id string, data entity.ProdukGudangCore) error {
	request := entity.ProdukGCoreToProdukGModel(data)

	tx := produkRepo.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}
