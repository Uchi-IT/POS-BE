package entity

type ProdukGudangRepositoryInterface interface {
	InputProduk(data ProdukGudangCore) (ProdukGudangCore, error)
	GetById(id string) (ProdukGudangCore, error)
	GetAllProduk(search, filter string) ([]ProdukGudangCore, error)
	UpdateProduk(id string, data ProdukGudangCore) error
	DeleteProduk(id string) error
}

type ProdukGudangServiceInterface interface {
	InputProduk(data ProdukGudangCore) (ProdukGudangCore, error)
	GetById(id string) (ProdukGudangCore, error)
	GetAllProduk(search, filter string) ([]ProdukGudangCore, error)
	UpdateProduk(id string, data ProdukGudangCore) error
	DeleteProduk(id string) error
}
