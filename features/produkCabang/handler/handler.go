package handler

import (
	"uchiiParfume/features/produkCabang/entity"

	"github.com/labstack/echo/v4"
)

type produkCHandler struct {
	produkCService entity.ProdukCabangServiceInterface
}

func NewProdukCHandler(produkC entity.ProdukCabangServiceInterface) *produkCHandler {
	return &produkCHandler{
		produkCService: produkC,
	}
}

func (produk *produkCHandler) InputProduk(e echo.Context) error {
	return nil
}

func (produk *produkCHandler) GetById(e echo.Context) error {
	return nil
}

func (produk *produkCHandler) GetAllProduk(e echo.Context) error {
	return nil
}

func (produk *produkCHandler) UpdateProduk(e echo.Context) error {
	return nil
}

func (produk *produkCHandler) DeleteProduk(e echo.Context) error {
	return nil
}
