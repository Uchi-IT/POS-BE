package handler

import (
	"uchiiParfume/features/notification/entity"

	"github.com/labstack/echo/v4"
)

type notifikasiHandler struct {
	notifikasiService entity.NotifikasiServiceInterface
}

func NewNotifikasiHandler(notifikasi entity.NotifikasiServiceInterface) *notifikasiHandler {
	return &notifikasiHandler{
		notifikasiService: notifikasi,
	}
}

func (notifikasi *notifikasiHandler) InputProduk(e echo.Context) error {
	return nil
}

func (notifikasi *notifikasiHandler) GetById(e echo.Context) error {
	return nil
}

func (notifikasi *notifikasiHandler) GetAllProduk(e echo.Context) error {
	return nil
}

func (notifikasi *notifikasiHandler) UpdateProduk(e echo.Context) error {
	return nil
}

func (notifikasi *notifikasiHandler) DeleteProduk(e echo.Context) error {
	return nil
}
