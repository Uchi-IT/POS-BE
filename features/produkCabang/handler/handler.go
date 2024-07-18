package handler

import (
	"net/http"
	"uchiiParfume/features/produkCabang/dto"
	"uchiiParfume/features/produkCabang/entity"
	middleware "uchiiParfume/utils/jwt"

	"github.com/google/uuid"
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

func (handler *produkCHandler) InputProduk(e echo.Context) error {
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

	input := new(dto.ProdukCabangRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.ProdukCabangCore{
		Foto:                input.Foto,
		NamaProduk:          input.NamaProduk,
		HargaJual:           input.HargaJual,
		SeratusMl:           input.SeratusMl,
		DuaratusMl:          input.DuaratusMl,
		DuaRatusLimaPuluhMl: input.DuaRatusLimaPuluhMl,
		Manual:              input.Manual,
	}

	row, errProduk := handler.produkCService.InputProduk(data)
	if errProduk != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create cabang",
			"error":   errProduk.Error(),
		})
	}

	resp := dto.ProdukCabangResponse{
		Id:                  row.Id,
		Foto:                row.Foto,
		NamaProduk:          row.NamaProduk,
		HargaJual:           row.HargaJual,
		SeratusMl:           row.SeratusMl,
		DuaratusMl:          row.DuaratusMl,
		DuaRatusLimaPuluhMl: row.DuaRatusLimaPuluhMl,
		Manual:              row.Manual,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create produk",
		"data":    resp,
	})
}

func (handler *produkCHandler) GetById(e echo.Context) error {
	idParamstr := e.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "produk not found",
		})
	}

	data, err := handler.produkCService.GetById(idParams.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get produk",
		})
	}

	response := dto.ProdukCabangResponse{
		Id:                  data.Id,
		Foto:                data.Foto,
		NamaProduk:          data.NamaProduk,
		HargaJual:           data.HargaJual,
		SeratusMl:           data.SeratusMl,
		DuaratusMl:          data.DuaratusMl,
		DuaRatusLimaPuluhMl: data.DuaRatusLimaPuluhMl,
		Manual:              data.Manual,
	}
	return e.JSON(http.StatusOK, map[string]any{
		"message": "get produk",
		"data":    response,
	})

}

func (handler *produkCHandler) GetAllProduk(e echo.Context) error {
	data, err := handler.produkCService.GetAllProduk()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all produk",
		})
	}

	dataList := []dto.ProdukCabangResponse{}
	for _, v := range data {
		result := dto.ProdukCabangResponse{
			Id:                  v.Id,
			Foto:                v.Foto,
			NamaProduk:          v.NamaProduk,
			HargaJual:           v.HargaJual,
			SeratusMl:           v.SeratusMl,
			DuaratusMl:          v.DuaratusMl,
			DuaRatusLimaPuluhMl: v.DuaRatusLimaPuluhMl,
			Manual:              v.Manual,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all produk",
		"data":    dataList,
	})
}

func (handler *produkCHandler) UpdateProduk(e echo.Context) error {
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

	data := new(dto.ProdukCabangRequest)
	if errBind := e.Bind(data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	produkData := entity.ProdukCabangCore{
		Foto:                data.Foto,
		NamaProduk:          data.NamaProduk,
		HargaJual:           data.HargaJual,
		SeratusMl:           data.SeratusMl,
		DuaratusMl:          data.DuaratusMl,
		DuaRatusLimaPuluhMl: data.DuaRatusLimaPuluhMl,
		Manual:              data.Manual,
	}	

	err := handler.produkCService.UpdateProduk(idParams, produkData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating produk",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "cabang updated successfully",
	})
}

func (handler *produkCHandler) DeleteProduk(e echo.Context) error {
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
	err := handler.produkCService.DeleteProduk(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting produk",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "produk deleted successfully",
	})
}
