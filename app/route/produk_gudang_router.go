package route

import (
	"uchiiParfume/features/produkGudang/handler"
	"uchiiParfume/features/produkGudang/repository"
	"uchiiParfume/features/produkGudang/service"
	m "uchiiParfume/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProdukGudangRoute(db *gorm.DB, e *echo.Group) {
	produkRepository := repository.NewProdukRepository(db)
	produkUseCase := service.NewProdukGService(produkRepository)
	produkController := handler.NewProdukGHandler(produkUseCase)

	produkGudang := e.Group("/produk-gudang")
	produkGudang.POST("", produkController.InputProduk, m.JWTMiddleware())
	produkGudang.GET("", produkController.GetAllProduk, m.JWTMiddleware())
	produkGudang.GET("/:id", produkController.GetById, m.JWTMiddleware())
	produkGudang.PUT("/:id", produkController.UpdateProduk, m.JWTMiddleware())
	produkGudang.DELETE("/:id", produkController.DeleteProduk, m.JWTMiddleware())
}
