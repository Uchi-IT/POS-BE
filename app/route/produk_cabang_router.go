package route

import (
	"uchiiParfume/features/produkCabang/handler"
	"uchiiParfume/features/produkCabang/repository"
	"uchiiParfume/features/produkCabang/service"
	m "uchiiParfume/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProdukCabangRoute(db *gorm.DB, e *echo.Group) {
	produkRepository := repository.NewProdukRepository(db)
	produkUseCase := service.NewProdukCService(produkRepository)
	producController := handler.NewProdukCHandler(produkUseCase)

	produkCabang := e.Group("/produk-cabang")
	produkCabang.POST("", producController.InputProduk, m.JWTMiddleware())
	produkCabang.GET("", producController.GetAllProduk, m.JWTMiddleware())
	produkCabang.GET("/:id", producController.GetById, m.JWTMiddleware())
	produkCabang.PUT("/:id", producController.UpdateProduk, m.JWTMiddleware())
	produkCabang.DELETE("/:id", producController.DeleteProduk, m.JWTMiddleware())
}
