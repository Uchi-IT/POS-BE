package service

import (
	"errors"
	"uchiiParfume/features/cabang/entity"
)

type cabangService struct {
	CabangRepository entity.CabangRepositoryInterface
}

func NewCabangService(cabang entity.CabangRepositoryInterface) entity.CabangServiceInterface {
	return &cabangService{
		CabangRepository: cabang,
	}
}

// CreateCabang implements entity.CabangServiceInterface.
func (cabangUC *cabangService) CreateCabang(data entity.CabangCore) (entity.CabangCore, error) {
	if data.Alamat == "" || data.NamaCabang == "" {
		return entity.CabangCore{}, errors.New("error, data can't be empty")
	}

	isExist, err := cabangUC.CabangRepository.IsNamaCabangExist(data.NamaCabang)
    if err != nil {
        return entity.CabangCore{}, err
    }
    if isExist {
        return entity.CabangCore{}, errors.New("nama cabang already exists")
    }

	errRegister, err := cabangUC.CabangRepository.CreateCabang(data)
	if err != nil {
		return entity.CabangCore{}, err
	}

	return errRegister, nil
}

// DeleteCabang implements entity.CabangServiceInterface.
func (cabangUC *cabangService) DeleteCabang(id string) error {
	if id == "" {
		return errors.New("insert cabang id")
	}

	errDelete := cabangUC.CabangRepository.DeleteCabang(id)
	if errDelete != nil {
		return errors.New("can't delete cabang")
	}

	return nil
}

// GetAllCabang implements entity.CabangServiceInterface.
func (cabangUC *cabangService) GetAllCabang() ([]entity.CabangCore, error) {
	cabang, err := cabangUC.CabangRepository.GetAllCabang()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return cabang, nil
}

// GetById implements entity.CabangServiceInterface.
func (cabangUC *cabangService) GetById(id string) (entity.CabangCore, error) {
	if id == "" {
		return entity.CabangCore{}, errors.New("cabang ID is required")
	}

	cabang, err := cabangUC.CabangRepository.GetById(id)
	if err != nil {
		return entity.CabangCore{}, err
	}

	return cabang, nil
}

// UpdateCabang implements entity.CabangServiceInterface.
func (cabangUC *cabangService) UpdateCabang(id string, data entity.CabangCore) error {
	if id == "" {
		return errors.New("id is not found")
	}

	if data.Alamat == "" || data.NamaCabang == "" {
		return errors.New("error, data can't be empty")
	}

	cabang, errGet := cabangUC.CabangRepository.GetById(id)
	if errGet != nil {
		return errGet
	}

	if cabang.Id == ""{
		return errors.New("cabang not found")
	}

	err := cabangUC.CabangRepository.UpdateCabang(id, data)
	if err != nil {
		return err
	}

	return nil
}
