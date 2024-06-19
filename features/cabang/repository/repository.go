package repository

import (
	"uchiiParfume/features/cabang/entity"

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
func (c *cabangRepository) CreateCabang(data entity.CabangCore) (entity.CabangCore, error) {
	panic("unimplemented")
}

// DeleteCabang implements entity.CabangRepositoryInterface.
func (c *cabangRepository) DeleteCabang(id string) error {
	panic("unimplemented")
}

// GetAllCabang implements entity.CabangRepositoryInterface.
func (c *cabangRepository) GetAllCabang() ([]entity.CabangCore, error) {
	panic("unimplemented")
}

// GetById implements entity.CabangRepositoryInterface.
func (c *cabangRepository) GetById(id string) (entity.CabangCore, error) {
	panic("unimplemented")
}

// UpdateCabang implements entity.CabangRepositoryInterface.
func (c *cabangRepository) UpdateCabang(id string, data entity.CabangCore) error {
	panic("unimplemented")
}