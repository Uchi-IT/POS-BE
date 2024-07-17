package handler

import (
	"net/http"
	dto "uchiiParfume/features/cabang/dto"
	"uchiiParfume/features/cabang/entity"
	middleware "uchiiParfume/utils/jwt"

	"github.com/google/uuid"
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

func (handler *cabangHandler) CreateCabang(e echo.Context) error {
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

	input := new(dto.CabangRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.CabangCore{
		Image:      input.Image,
		NamaCabang: input.NamaCabang,
		Alamat:     input.Alamat,
	}

	row, errCabang := handler.cabangService.CreateCabang(data)
	if errCabang != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create cabang",
			"error":   errCabang.Error(),
		})
	}

	resp := dto.CabangResponse{
		Id:         row.Id,
		Image:      row.Image,
		NamaCabang: row.NamaCabang,
		Alamat:     row.Alamat,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create cabang",
		"data":    resp,
	})
}

func (handler *cabangHandler) GetById(e echo.Context) error {
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

	idParamstr := e.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "cabang not found",
		})
	}

	data, err := handler.cabangService.GetById(idParams.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get cabang",
		})
	}

	response := dto.CabangResponse{
		Id:         data.Id,
		Image:      data.Image,
		NamaCabang: data.NamaCabang,
		Alamat:     data.Alamat,
	}
	return e.JSON(http.StatusOK, map[string]any{
		"message": "get cabang",
		"data":    response,
	})
}

func (handler *cabangHandler) GetAllCabang(e echo.Context) error {
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

	data, err := handler.cabangService.GetAllCabang()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all cabang",
		})
	}

	dataList := []dto.CabangResponse{}
	for _, v := range data {
		result := dto.CabangResponse{
			Id:         v.Id,
			Image:      v.Image,
			NamaCabang: v.NamaCabang,
			Alamat:     v.Alamat,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all cabang",
		"data":    dataList,
	})
}

func (handler *cabangHandler) UpdateCabang(e echo.Context) error {
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

	data := new(dto.CabangRequest)
	if errBind := e.Bind(data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	cabangData := entity.CabangCore{
		Image:      data.Image,
		NamaCabang: data.NamaCabang,
		Alamat:     data.Alamat,
	}	

	err := handler.cabangService.UpdateCabang(idParams, cabangData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating cabang",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "cabang updated successfully",
	})
}

func (handler *cabangHandler) DeleteCabang(e echo.Context) error {
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
	err := handler.cabangService.DeleteCabang(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting cabang",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Cabang deleted successfully",
	})
}
