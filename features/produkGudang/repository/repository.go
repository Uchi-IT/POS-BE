package repository

import (
	"uchiiParfume/features/produkGudang/entity"

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
func (p *produkGRepository) DeleteProduk(id string) error {
	panic("unimplemented")
}

// GetAllProduk implements entity.ProdukGudangRepositoryInterface.
func (p *produkGRepository) GetAllProduk() ([]entity.ProdukGudangCore, error) {
	panic("unimplemented")
}

// GetById implements entity.ProdukGudangRepositoryInterface.
func (p *produkGRepository) GetById(id string) (entity.ProdukGudangCore, error) {
	panic("unimplemented")
}

// InputProduk implements entity.ProdukGudangRepositoryInterface.
func (p *produkGRepository) InputProduk(data entity.ProdukGudangCore) (entity.ProdukGudangCore, error) {
	panic("unimplemented")
}

// UpdateProduk implements entity.ProdukGudangRepositoryInterface.
func (p *produkGRepository) UpdateProduk(id string, data entity.ProdukGudangCore) error {
	panic("unimplemented")
}