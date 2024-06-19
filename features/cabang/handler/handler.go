package handler

import (
	"uchiiParfume/features/cabang/entity"

	"github.com/labstack/echo/v4"
)

type cabangHandler struct {
	cabangService entity.CabangServiceInterface
}

func NewCabangHandler(cabang entity.CabangServiceInterface) *cabangHandler {
	return &cabangHandler{
		cabangService: cabang,
	}
}

func (cabang *cabangHandler) CreateCabang(e echo.Context) error {
	return nil
}

func (cabang *cabangHandler) GetById(e echo.Context) error {
	return nil
}

func (cabang *cabangHandler) GetAllCabang(e echo.Context) error {
	return nil
}

func (cabang *cabangHandler) UpdateCabang(e echo.Context) error {
	return nil
}

func (cabang *cabangHandler) DeleteCabang(e echo.Context) error {
	return nil
}
