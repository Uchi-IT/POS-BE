package service

import "uchiiParfume/features/cabang/entity"

type cabangService struct {
	CabangRepository entity.CabangRepositoryInterface
}

func NewCabangService(cabang entity.CabangRepositoryInterface) entity.CabangServiceInterface {
	return &cabangService{
		CabangRepository: cabang,
	}
}

// CreateCabang implements entity.CabangServiceInterface.
func (c *cabangService) CreateCabang(data entity.CabangCore) (entity.CabangCore, error) {
	panic("unimplemented")
}

// DeleteCabang implements entity.CabangServiceInterface.
func (c *cabangService) DeleteCabang(id string) error {
	panic("unimplemented")
}

// GetAllCabang implements entity.CabangServiceInterface.
func (c *cabangService) GetAllCabang() ([]entity.CabangCore, error) {
	panic("unimplemented")
}

// GetById implements entity.CabangServiceInterface.
func (c *cabangService) GetById(id string) (entity.CabangCore, error) {
	panic("unimplemented")
}

// UpdateCabang implements entity.CabangServiceInterface.
func (c *cabangService) UpdateCabang(id string, data entity.CabangCore) error {
	panic("unimplemented")
}