package service

import (
	"errors"
	cabang "uchiiParfume/features/cabang/entity"
	"uchiiParfume/features/produkCabang/entity"
	produk "uchiiParfume/features/produkGudang/entity"
	"uchiiParfume/utils/constanta"
	"uchiiParfume/utils/validation"
)

type produkCService struct {
	ProdukRepository       entity.ProdukCabangRepositoryInterface
	CabangRepository       cabang.CabangRepositoryInterface
	ProdukGudangRepository produk.ProdukGudangRepositoryInterface
}

func NewProdukCService(produkC entity.ProdukCabangRepositoryInterface, cabang cabang.CabangRepositoryInterface, produkG produk.ProdukGudangRepositoryInterface) entity.ProdukCabangServiceInterface {
	return &produkCService{
		ProdukRepository:       produkC,
		CabangRepository:       cabang,
		ProdukGudangRepository: produkG,
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
func (produkUC *produkCService) GetAllProduk(search, filter string) ([]entity.ProdukCabangCore, error) {

	if filter != "asc" && filter != "desc" {
		filterData, errEqual := validation.CheckEqualData(filter, constanta.PRODUCT_CABANG)
		if errEqual != nil {
			return []entity.ProdukCabangCore{}, errors.New("error : filter tidak valid")
		}
		filter = filterData
	}

	cabang, err := produkUC.ProdukRepository.GetAllProduk(search, filter)
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
func (produkUC *produkCService) InputProduk(data []entity.ProdukCabangCore, riwayat entity.RiwayatProdukCabangCore) ([]entity.ProdukCabangCore, error) {
	var results []entity.ProdukCabangCore

	for _, produk := range data {
		_, errc := produkUC.CabangRepository.GetById(produk.CabangId)
		if errc != nil {
			return nil, errors.New("cabang not found")
		}

		dataGudang, errg := produkUC.ProdukGudangRepository.GetById(produk.ProdukId)
		if errg != nil {
			return nil, errors.New("produk not found")
		}

		produk.NamaProduk = dataGudang.NamaProduk

		if produk.NamaProduk == "" {
			return nil, errors.New("nama produk can't empty")
		}

		if produk.HargaJual < 0 {
			return nil, errors.New("harga produk can't be less than 0")
		}

		count100 := produk.SeratusMl * 100
		count200 := produk.DuaratusMl * 200
		count250 := produk.DuaRatusLimaPuluhMl * 250

		countTotal := count100 + count200 + count250 + produk.Manual
		produk.Total = countTotal

		// Call repository to insert each product
		errInput, err := produkUC.ProdukRepository.InputProduk(produk)
		if err != nil {
			return nil, err
		}

		// Menambahkan detail ke data riwayat
		riwayat.ProdukDetail = append(riwayat.ProdukDetail, entity.RiwayatProdukCabangItemCore{
			ProdukCabangId:      dataGudang.Id,
			NamaProduk:          produk.NamaProduk,
			SeratusMl:           produk.SeratusMl,
			DuaratusMl:          produk.DuaratusMl,
			DuaRatusLimaPuluhMl: produk.DuaRatusLimaPuluhMl,
			Manual:              produk.Manual,
		})

		results = append(results, errInput)
	}

	riwayat.TotalStok = calculateTotalStok(riwayat.ProdukDetail) // Function to calculate total stock
	_, errRiwayat := produkUC.ProdukRepository.InputRiwayat(riwayat)
	if errRiwayat != nil {
		return nil, errRiwayat
	}

	return results, nil
}

// UpdateProduk implements entity.ProdukCabangServiceInterface.
func (produkUC *produkCService) UpdateProduk(id string, data entity.ProdukCabangCore) error {
	if data.NamaProduk == "" {
		return errors.New("nama produk can't empty")
	}

	if data.HargaJual < 0 {
		return errors.New("harga produk can't less then 0")
	}

	_, errc := produkUC.CabangRepository.GetById(data.CabangId)
	if errc != nil {
		return errors.New("cabang not found")
	}

	_, errg := produkUC.ProdukGudangRepository.GetById(data.ProdukId)
	if errg != nil {
		return errors.New("produk not found")
	}

	_, errGet := produkUC.ProdukRepository.GetById(id)
	if errGet != nil {
		return errGet
	}

	err := produkUC.ProdukRepository.UpdateProduk(id, data)
	if err != nil {
		return err
	}

	return nil
}

func calculateTotalStok(items []entity.RiwayatProdukCabangItemCore) int {
	total := 0
	for _, item := range items {
		total += item.SeratusMl*100 + item.DuaratusMl*200 + item.DuaRatusLimaPuluhMl*250 + item.Manual
	}
	return total
}

// GetAllRiwayat implements entity.ProdukCabangServiceInterface.
func (produkUC *produkCService) GetAllRiwayat() ([]entity.RiwayatProdukCabangCore, error) {
	riwayat, err := produkUC.ProdukRepository.GetAllRiwayat()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return riwayat, nil
}