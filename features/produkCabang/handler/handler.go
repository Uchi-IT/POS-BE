package handler

import (
	"net/http"
	ce "uchiiParfume/features/cabang/entity"
	"uchiiParfume/features/produkCabang/dto"
	"uchiiParfume/features/produkCabang/entity"
	ue "uchiiParfume/features/users/entity"
	middleware "uchiiParfume/utils/jwt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type produkCHandler struct {
	produkCService entity.ProdukCabangServiceInterface
	userService    ue.UsersServiceInterface
	cabangService  ce.CabangServiceInterface
}

func NewProdukCHandler(produkC entity.ProdukCabangServiceInterface, user ue.UsersServiceInterface, cabang ce.CabangServiceInterface) *produkCHandler {
	return &produkCHandler{
		produkCService: produkC,
		userService:    user,
		cabangService:  cabang,
	}
}

func (handler *produkCHandler) InputProduk(e echo.Context) error {
    var namaCabang string

    // Extract token and check role
    userId, role, err := middleware.ExtractToken(e)
    if err != nil {
        return e.JSON(http.StatusBadRequest, map[string]any{
            "message": err.Error(),
        })
    }

    // Only admin can access this route
    if role != "admin" {
        return e.JSON(http.StatusForbidden, map[string]any{
            "message": "access denied",
        })
    }

    // Mengisi data nama admin dan cabang
    userData, err := handler.userService.GetById(userId)
    if err != nil {
        return e.JSON(http.StatusBadRequest, map[string]any{
            "message": "error get user data",
        })
    }

    // Bind the request data
    var input dto.ProdukCabangWithNote
    errBind := e.Bind(&input)
    if errBind != nil {
        return e.JSON(http.StatusBadRequest, map[string]any{
            "message": "error bind data",
            "error":   errBind.Error(),
        })
    }

    // Prepare slice for ProdukCabangCore
    var data []entity.ProdukCabangCore
    for _, produk := range input.Produk {
        data = append(data, entity.ProdukCabangCore{
            Foto:                produk.Foto,
            HargaJual:           produk.HargaJual,
            SeratusMl:           produk.SeratusMl,
            DuaratusMl:          produk.DuaratusMl,
            Manual:              produk.Manual,
            CabangId:            produk.CabangId,
            ProdukId:            produk.ProdukId,
        })

        // Get cabang data based on CabangId
        cabangData, err := handler.cabangService.GetById(produk.CabangId)
        if err != nil {
            return e.JSON(http.StatusBadRequest, map[string]any{
                "message": "error get cabang data",
            })
        }

        namaCabang = cabangData.NamaCabang
    }

    // Convert riwayat request to RiwayatProdukCabang entity
    riwayat := entity.RiwayatProdukCabangCore{
        NamaAdmin:  userData.Nama,
        NamaCabang: namaCabang,
        Catatan:    input.Catatan.Catatan, // Set catatan from request
    }

    // Call the service to insert the products
    _, errProduk := handler.produkCService.InputProduk(data, riwayat)
    if errProduk != nil {
        return e.JSON(http.StatusBadRequest, map[string]any{
            "message": "error creating products",
            "error":   errProduk.Error(),
        })
    }

    // Return success response
    return e.JSON(http.StatusOK, map[string]any{
        "message": "success create products",
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
		Total:               data.Total,
	}
	return e.JSON(http.StatusOK, map[string]any{
		"message": "get produk",
		"data":    response,
	})
}

func (handler *produkCHandler) GetAllProduk(e echo.Context) error {
	search := e.QueryParam("search")
	sort := e.QueryParam("sort")
	data, err := handler.produkCService.GetAllProduk(search, sort)
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
			Total:               v.Total,
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
		HargaJual:           data.HargaJual,
		SeratusMl:           data.SeratusMl,
		DuaratusMl:          data.DuaratusMl,
		DuaRatusLimaPuluhMl: data.DuaRatusLimaPuluhMl,
		Manual:              data.Manual,
		ProdukId:            data.ProdukId,
		CabangId:            data.CabangId,
	}

	err := handler.produkCService.UpdateProduk(idParams, produkData)
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

func (handler *produkCHandler) GetAllRiwayat(e echo.Context) error {
	data, err := handler.produkCService.GetAllRiwayat()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all transaction",
		})
	}

	dataList := dto.ListTransactionCoreToListTransactionResponse(data)

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all transaction",
		"data":    dataList,
	})
}