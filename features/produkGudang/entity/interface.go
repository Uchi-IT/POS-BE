package entity

type ProdukGudangRepositoryInterface interface {
	InputProduk(data ProdukGudangCore) (ProdukGudangCore, error)
	GetById(id string) (ProdukGudangCore, error)
	GetAllProduk() ([]ProdukGudangCore, error)
	UpdateProduk(id string, data ProdukGudangCore) error
	DeleteProduk(id string) error
}

type ProdukGudangServiceInterface interface {
	InputProduk(data ProdukGudangCore) (ProdukGudangCore, error)
	GetById(id string) (ProdukGudangCore, error)
	GetAllProduk() ([]ProdukGudangCore, error)
	UpdateProduk(id string, data ProdukGudangCore) error
	DeleteProduk(id string) error
}
