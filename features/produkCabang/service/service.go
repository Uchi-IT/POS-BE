package service

import (
	"errors"
	"uchiiParfume/features/produkCabang/entity"
)

type produkCService struct {
	ProdukRepository entity.ProdukCabangRepositoryInterface
}

func NewProdukCService(produkC entity.ProdukCabangRepositoryInterface) entity.ProdukCabangServiceInterface {
	return &produkCService{
		ProdukRepository: produkC,
	}
}

// DeleteProduk implements entity.ProdukCabangServiceInterface.
func (produkUC *produkCService) DeleteProduk(id string) error {
	if id == "" {
		return errors.New("insert produk id")
	}

	errDelete := produkUC.ProdukRepository.DeleteProduk(id)
	if errDelete != nil {
		return errors.New("can't delete produk")
	}

	return nil
}

// GetAllProduk implements entity.ProdukCabangServiceInterface.
func (produkUC *produkCService) GetAllProduk() ([]entity.ProdukCabangCore, error) {
	cabang, err := produkUC.ProdukRepository.GetAllProduk()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return cabang, nil
}

// GetById implements entity.ProdukCabangServiceInterface.
func (produkUC *produkCService) GetById(id string) (entity.ProdukCabangCore, error) {
	if id == "" {
		return entity.ProdukCabangCore{}, errors.New("produk ID is required")
	}

	cabang, err := produkUC.ProdukRepository.GetById(id)
	if err != nil {
		return entity.ProdukCabangCore{}, err
	}

	return cabang, nil
}

// InputProduk implements entity.ProdukCabangServiceInterface.
func (produkUC *produkCService) InputProduk(data entity.ProdukCabangCore) (entity.ProdukCabangCore, error) {
	if data.NamaProduk == "" {
		return entity.ProdukCabangCore{}, errors.New("nama produk can't empty")
	}

	if data.HargaJual < 0{
		return entity.ProdukCabangCore{}, errors.New("harga produk can't less then 0")
	}

	errInput, err := produkUC.ProdukRepository.InputProduk(data)
	if err != nil{
		return entity.ProdukCabangCore{}, err
	}

	return errInput, nil
}

// UpdateProduk implements entity.ProdukCabangServiceInterface.
func (produkUC *produkCService) UpdateProduk(id string, data entity.ProdukCabangCore) error {
	if data.NamaProduk == "" {
		return errors.New("nama produk can't empty")
	}

	if data.HargaJual < 0{
		return errors.New("harga produk can't less then 0")
	}

	_, errGet := produkUC.ProdukRepository.GetById(id)
	if errGet != nil {
		return errGet
	}

	err := produkUC.ProdukRepository.UpdateProduk(id,data)
	if err != nil {
		return err
	}

	return nil
}
