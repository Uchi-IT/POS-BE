package service

import "uchiiParfume/features/produkCabang/entity"

type produkCService struct {
	ProdukRepository entity.ProdukCabangRepositoryInterface
}

func NewProdukCService(produkC entity.ProdukCabangRepositoryInterface) entity.ProdukCabangServiceInterface {
	return &produkCService{
		ProdukRepository: produkC,
	}
}

// DeleteProduk implements entity.ProdukCabangServiceInterface.
func (p *produkCService) DeleteProduk(id string) error {
	panic("unimplemented")
}

// GetAllProduk implements entity.ProdukCabangServiceInterface.
func (p *produkCService) GetAllProduk() ([]entity.ProdukCabangCore, error) {
	panic("unimplemented")
}

// GetById implements entity.ProdukCabangServiceInterface.
func (p *produkCService) GetById(id string) (entity.ProdukCabangCore, error) {
	panic("unimplemented")
}

// InputProduk implements entity.ProdukCabangServiceInterface.
func (p *produkCService) InputProduk(data entity.ProdukCabangCore) (entity.ProdukCabangCore, error) {
	panic("unimplemented")
}

// UpdateProduk implements entity.ProdukCabangServiceInterface.
func (p *produkCService) UpdateProduk(id string, data entity.ProdukCabangCore) error {
	panic("unimplemented")
}
