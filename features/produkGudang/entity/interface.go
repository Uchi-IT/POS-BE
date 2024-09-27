package entity

import "uchiiParfume/utils/pagination"

type ProdukGudangRepositoryInterface interface {
	InputProduk(data ProdukGudangCore) (ProdukGudangCore, error)
	GetById(id string) (ProdukGudangCore, error)
	GetAllProduk(page, limit int, search, filter string) ([]ProdukGudangCore, pagination.PageInfo, int, error)
	UpdateProduk(id string, data ProdukGudangCore) error
	DeleteProduk(id string) error
}

type ProdukGudangServiceInterface interface {
	InputProduk(data ProdukGudangCore) (ProdukGudangCore, error)
	GetById(id string) (ProdukGudangCore, error)
	GetAllProduk(page, limit int, search, filter string) ([]ProdukGudangCore, pagination.PageInfo, int, error)
	UpdateProduk(id string, data ProdukGudangCore) error
	DeleteProduk(id string) error
}
