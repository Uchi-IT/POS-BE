package entity

type CabangRepositoryInterface interface {
	CreateCabang(data CabangCore) (CabangCore, error)
	GetById(id string) (CabangCore, error)
	GetAllCabang() ([]CabangCore, error)
	UpdateCabang(id string, data CabangCore) error
	DeleteCabang(id string) error
}

type CabangServiceInterface interface{
	CreateCabang(data CabangCore) (CabangCore, error)
	GetById(id string) (CabangCore, error)
	GetAllCabang() ([]CabangCore, error)
	UpdateCabang(id string, data CabangCore) error
	DeleteCabang(id string) error
}