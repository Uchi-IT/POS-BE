package handler

import (
	"uchiiParfume/features/produkGudang/entity"

	"github.com/labstack/echo/v4"
)

type produkGHandler struct {
	produkGService entity.ProdukGudangServiceInterface
}

func NewProdukGHandler(produkG entity.ProdukGudangServiceInterface) *produkGHandler {
	return &produkGHandler{
		produkGService: produkG,
	}
}

func (produk *produkGHandler) InputProduk(e echo.Context) error {
	return nil
}

func (produk *produkGHandler) GetById(e echo.Context) error {
	return nil
}

func (produk *produkGHandler) GetAllProduk(e echo.Context) error {
	return nil
}

func (produk *produkGHandler) UpdateProduk(e echo.Context) error {
	return nil
}

func (produk *produkGHandler) DeleteProduk(e echo.Context) error {
	return nil
}
