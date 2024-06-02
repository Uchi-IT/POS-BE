package entity

type ProdukCabangRepositoryInterface interface {
	InputProduk(data ProdukCabangCore) (ProdukCabangCore, error)
	GetById(id string) (ProdukCabangCore, error)
	GetAllProduk() ([]ProdukCabangCore, error)
	UpdateProduk(id string, data ProdukCabangCore) error
	DeleteProduk(id string) error
}

type ProdukCabangServiceInterface interface {
	InputProduk(data ProdukCabangCore) (ProdukCabangCore, error)
	GetById(id string) (ProdukCabangCore, error)
	GetAllProduk() ([]ProdukCabangCore, error)
	UpdateProduk(id string, data ProdukCabangCore) error
	DeleteProduk(id string) error
}
