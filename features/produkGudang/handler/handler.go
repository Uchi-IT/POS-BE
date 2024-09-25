package handler

import (
	"net/http"
	"uchiiParfume/features/produkGudang/dto"
	"uchiiParfume/features/produkGudang/entity"
	middleware "uchiiParfume/utils/jwt"

	"github.com/google/uuid"
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

func (handler *produkGHandler) InputProduk(e echo.Context) error {
	_, role, err := middleware.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	input := new(dto.ProdukGudangRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.ProdukGudangCore{
		Foto:       input.Foto,
		NamaProduk: input.NamaProduk,
		Stok:       input.Stok,
		HargaJual:  input.HargaJual,
		HargaBeli:  input.HargaBeli,
	}

	row, errProduk := handler.produkGService.InputProduk(data)
	if errProduk != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create produk",
			"error":   errProduk.Error(),
		})
	}

	resp := dto.ProdukGudangResponse{
		Id:         row.Id,
		Foto:       row.Foto,
		NamaProduk: row.NamaProduk,
		Stok:       row.Stok,
		HargaJual:  row.HargaJual,
		HargaBeli:  row.HargaBeli,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create produk",
		"data":    resp,
	})
}

func (handler *produkGHandler) GetById(e echo.Context) error {
	_, role, errRole := middleware.ExtractToken(e)
	if errRole != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": errRole.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParamstr := e.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "produk not found",
		})
	}

	data, err := handler.produkGService.GetById(idParams.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get produk",
		})
	}

	response := dto.ProdukGudangResponse{
		Id:         data.Id,
		Foto:       data.Foto,
		NamaProduk: data.NamaProduk,
		Stok:       data.Stok,
		HargaJual:  data.HargaJual,
		HargaBeli:  data.HargaBeli,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get produk",
		"data":    response,
	})
}

func (handler *produkGHandler) GetAllProduk(e echo.Context) error {
	_, role, errRole := middleware.ExtractToken(e)
	if errRole != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": errRole.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	search := e.QueryParam("search")
	filter := e.QueryParam("filter")

	data, err := handler.produkGService.GetAllProduk(search, filter)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all produk",
		})
	}

	dataList := []dto.ProdukGudangResponse{}
	for _, v := range data {
		result := dto.ProdukGudangResponse{
			Id:         v.Id,
			Foto:       v.Foto,
			NamaProduk: v.NamaProduk,
			Stok:       v.Stok,
			HargaJual:  v.HargaJual,
			HargaBeli:  v.HargaBeli,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all produk",
		"data":    dataList,
	})
}

func (handler *produkGHandler) UpdateProduk(e echo.Context) error {
	_, role, errRole := middleware.ExtractToken(e)
	if errRole != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": errRole.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")

	data := new(dto.ProdukGudangRequest)
	if errBind := e.Bind(data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	produkData := entity.ProdukGudangCore{
		Foto:       data.Foto,
		NamaProduk: data.NamaProduk,
		Stok:       data.Stok,
		HargaJual:  data.HargaJual,
		HargaBeli:  data.HargaBeli,
	}

	err := handler.produkGService.UpdateProduk(idParams, produkData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating produk",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "produk updated successfully",
	})
}

func (handler *produkGHandler) DeleteProduk(e echo.Context) error {
	_, role, errRole := middleware.ExtractToken(e)
	if errRole != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": errRole.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")
	err := handler.produkGService.DeleteProduk(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting produk",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "produk deleted successfully",
	})
}
