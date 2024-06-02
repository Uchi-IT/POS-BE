package repository

import (
	"uchiiParfume/features/produkCabang/entity"

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
func (p *produkCRepository) DeleteProduk(id string) error {
	panic("unimplemented")
}

// GetAllProduk implements entity.ProdukCabangRepositoryInterface.
func (p *produkCRepository) GetAllProduk() ([]entity.ProdukCabangCore, error) {
	panic("unimplemented")
}

// GetById implements entity.ProdukCabangRepositoryInterface.
func (p *produkCRepository) GetById(id string) (entity.ProdukCabangCore, error) {
	panic("unimplemented")
}

// InputProduk implements entity.ProdukCabangRepositoryInterface.
func (p *produkCRepository) InputProduk(data entity.ProdukCabangCore) (entity.ProdukCabangCore, error) {
	panic("unimplemented")
}

// UpdateProduk implements entity.ProdukCabangRepositoryInterface.
func (p *produkCRepository) UpdateProduk(id string, data entity.ProdukCabangCore) error {
	panic("unimplemented")
}