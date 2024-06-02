package service

import "uchiiParfume/features/produkGudang/entity"

type produkGService struct {
	ProdukRepository entity.ProdukGudangRepositoryInterface
}

func NewProdukGService(produkG entity.ProdukGudangRepositoryInterface) entity.ProdukGudangServiceInterface {
	return &produkGService{
		ProdukRepository: produkG,
	}
}

// DeleteProduk implements entity.ProdukGudangServiceInterface.
func (p *produkGService) DeleteProduk(id string) error {
	panic("unimplemented")
}

// GetAllProduk implements entity.ProdukGudangServiceInterface.
func (p *produkGService) GetAllProduk() ([]entity.ProdukGudangCore, error) {
	panic("unimplemented")
}

// GetById implements entity.ProdukGudangServiceInterface.
func (p *produkGService) GetById(id string) (entity.ProdukGudangCore, error) {
	panic("unimplemented")
}

// InputProduk implements entity.ProdukGudangServiceInterface.
func (p *produkGService) InputProduk(data entity.ProdukGudangCore) (entity.ProdukGudangCore, error) {
	panic("unimplemented")
}

// UpdateProduk implements entity.ProdukGudangServiceInterface.
func (p *produkGService) UpdateProduk(id string, data entity.ProdukGudangCore) error {
	panic("unimplemented")
}