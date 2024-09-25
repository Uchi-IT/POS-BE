package service

import (
	"errors"
	"uchiiParfume/features/produkGudang/entity"
)

type produkGService struct {
	ProdukRepository entity.ProdukGudangRepositoryInterface
}

func NewProdukGService(produkG entity.ProdukGudangRepositoryInterface) entity.ProdukGudangServiceInterface {
	return &produkGService{
		ProdukRepository: produkG,
	}
}

// DeleteProduk implements entity.ProdukGudangServiceInterface.
func (produkUC *produkGService) DeleteProduk(id string) error {
	if id == "" {
		return errors.New("insert produk id")
	}

	errDelete := produkUC.ProdukRepository.DeleteProduk(id)
	if errDelete != nil {
		return errors.New("can't delete produk")
	}

	return nil
}

// GetAllProduk implements entity.ProdukGudangServiceInterface.
func (produkUC *produkGService) GetAllProduk(search, filter string) ([]entity.ProdukGudangCore, error) {
	cabang, err := produkUC.ProdukRepository.GetAllProduk(search, filter)
	if err != nil {
		return nil, errors.New("error get data")
	}

	return cabang, nil
}

// GetById implements entity.ProdukGudangServiceInterface.
func (produkUC *produkGService) GetById(id string) (entity.ProdukGudangCore, error) {
	if id == "" {
		return entity.ProdukGudangCore{}, errors.New("produk ID is required")
	}

	cabang, err := produkUC.ProdukRepository.GetById(id)
	if err != nil {
		return entity.ProdukGudangCore{}, err
	}

	return cabang, nil
}

// InputProduk implements entity.ProdukGudangServiceInterface.
func (produkUC *produkGService) InputProduk(data entity.ProdukGudangCore) (entity.ProdukGudangCore, error) {
	if data.NamaProduk == "" {
		return entity.ProdukGudangCore{}, errors.New("nama produk can't empty")
	}

	if data.Stok < 0 {
		return entity.ProdukGudangCore{}, errors.New("stok can't less then 0")
	}

	if data.HargaJual < 0 || data.HargaBeli < 0{
		return entity.ProdukGudangCore{}, errors.New("harga can't less then 0")
	}

	errInput, err := produkUC.ProdukRepository.InputProduk(data)
	if err != nil{
		return entity.ProdukGudangCore{}, err
	}

	return errInput, nil
}

// UpdateProduk implements entity.ProdukGudangServiceInterface.
func (produkUC *produkGService) UpdateProduk(id string, data entity.ProdukGudangCore) error {
	if data.NamaProduk == "" {
		return errors.New("nama produk can't empty")
	}

	if data.Stok < 0 {
		return errors.New("stok can't less then 0")
	}

	if data.HargaJual < 0 || data.HargaBeli < 0{
		return errors.New("harga can't less then 0")
	}

	err := produkUC.ProdukRepository.UpdateProduk(id,data)
	if err != nil{
		return err
	}

	return nil
}
