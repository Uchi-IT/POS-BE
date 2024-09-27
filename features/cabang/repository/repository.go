package repository

import (
	"errors"
	"uchiiParfume/features/cabang/entity"
	"uchiiParfume/features/cabang/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type cabangRepository struct {
	db *gorm.DB
}

func NewCabangRepository(db *gorm.DB) entity.CabangRepositoryInterface {
	return &cabangRepository{
		db: db,
	}
}

// CreateCabang implements entity.CabangRepositoryInterface.
func (cabangRepo *cabangRepository) CreateCabang(data entity.CabangCore) (entity.CabangCore, error) {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return entity.CabangCore{}, UUIDerr
	}

	input := entity.CabangCoreToCabangModel(data)
	input.Id = newUUID.String()

	errCabang := cabangRepo.db.Save(&input)
	if errCabang.Error != nil {
		return entity.CabangCore{}, errCabang.Error
	}

	var resp = entity.CabangModelToCabangCore(input)

	return resp, nil
}

// DeleteCabang implements entity.CabangRepositoryInterface.
func (cabangRepo *cabangRepository) DeleteCabang(id string) error {
	var chekcId model.Cabang

	errData := cabangRepo.db.Where("id = ?", id).Delete(&chekcId)
	if errData != nil {
		return errData.Error
	}

	return nil
}

// GetAllCabang implements entity.CabangRepositoryInterface.
func (cabangRepo *cabangRepository) GetAllCabang(search, filter string) ([]entity.CabangCore, error) {
	var dataCabang []model.Cabang

	// Initialize query
	query := cabangRepo.db.Model(&model.Cabang{})

	// Add search functionality
	if search != "" {
		query = query.Where("cabangs.nama_cabang ILIKE ?", "%"+search+"%") // Use LIKE for partial search
	}

	// Add sorting based on filter
	if filter == "nama_asc" {
		query = query.Order("cabangs.nama_cabang ASC") // Correct column name
	} else if filter == "nama_desc" {
		query = query.Order("cabangs.nama_cabang DESC") // Correct column name
	} else if filter == "lates_update" {
		query = query.Order("cabangs.updated_at DESC") // Sort by last updated
	} else {
		query = query.Order("cabangs.created_at DESC") // Default sorting by created date
	}

	// Execute the query
	tx := query.Find(&dataCabang)
	if tx.Error != nil { // Correct error checking
		return nil, tx.Error
	}

	// Map data from model to core
	mapData := entity.ListCabangModelToCabangCore(dataCabang)
	return mapData, nil
}


// GetById implements entity.CabangRepositoryInterface.
func (cabangRepo *cabangRepository) GetById(id string) (entity.CabangCore, error) {
	var data model.Cabang
	errData := cabangRepo.db.Where("id=?", id).First(&data).Error
	if errData != nil {
		return entity.CabangCore{}, errData
	}

	resp := entity.CabangModelToCabangCore(data)

	return resp, nil
}

// UpdateCabang implements entity.CabangRepositoryInterface.
func (cabangRepo *cabangRepository) UpdateCabang(id string, data entity.CabangCore) error {
	request := entity.CabangCoreToCabangModel(data)

	tx := cabangRepo.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("cabang not found")
	}

	return nil
}

// IsNamaCabangExist implements entity.CabangRepositoryInterface.
func (cabangRepo *cabangRepository) IsNamaCabangExist(namaCabang string) (bool, error) {
	var count int64
	err := cabangRepo.db.Model(&model.Cabang{}).Where("nama_cabang = ?", namaCabang).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
