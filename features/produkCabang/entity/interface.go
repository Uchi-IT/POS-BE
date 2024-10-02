	package entity

type ProdukCabangRepositoryInterface interface {
	InputProduk(data ProdukCabangCore) (ProdukCabangCore, error)
	GetById(id string) (ProdukCabangCore, error)
	GetAllProduk(search, filter string) ([]ProdukCabangCore, error)
	UpdateProduk(id string, data ProdukCabangCore) error
	DeleteProduk(id string) error

	InputRiwayat(data RiwayatProdukCabangCore) (RiwayatProdukCabangCore, error)
	GetAllRiwayat() ([]RiwayatProdukCabangCore, error)
}

type ProdukCabangServiceInterface interface {
	InputProduk(data []ProdukCabangCore, riwayat RiwayatProdukCabangCore) ([]ProdukCabangCore, error)
	GetById(id string) (ProdukCabangCore, error)
	GetAllProduk(search, filter string) ([]ProdukCabangCore, error)
	UpdateProduk(id string, data ProdukCabangCore) error
	DeleteProduk(id string) error

	GetAllRiwayat() ([]RiwayatProdukCabangCore, error)
}
