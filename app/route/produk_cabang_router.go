package route

import (
	"uchiiParfume/features/produkCabang/handler"
	"uchiiParfume/features/produkCabang/repository"
	produkG "uchiiParfume/features/produkGudang/repository"
	"uchiiParfume/features/produkCabang/service"
	rc "uchiiParfume/features/cabang/repository"
	sc "uchiiParfume/features/cabang/service"
	ru "uchiiParfume/features/users/repository"
	su "uchiiParfume/features/users/service"
	m "uchiiParfume/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProdukCabangRoute(db *gorm.DB, e *echo.Group) {

	cabangRepository := rc.NewCabangRepository(db)
	cabangUseCase := sc.NewCabangService(cabangRepository)

	userRepository := ru.NewUserRepository(db,cabangRepository)
	userUseCase := su.NewUserService(userRepository)

	produkGRepository := produkG.NewProdukRepository(db)

	produkRepository := repository.NewProdukRepository(db)
	produkUseCase := service.NewProdukCService(produkRepository,cabangRepository,produkGRepository)
	producController := handler.NewProdukCHandler(produkUseCase, userUseCase, cabangUseCase)

	produkCabang := e.Group("/produk-cabang")
	produkCabang.POST("", producController.InputProduk, m.JWTMiddleware())
	produkCabang.GET("", producController.GetAllProduk, m.JWTMiddleware())
	produkCabang.GET("/:id", producController.GetById, m.JWTMiddleware())
	produkCabang.PUT("/:id", producController.UpdateProduk, m.JWTMiddleware())
	produkCabang.DELETE("/:id", producController.DeleteProduk, m.JWTMiddleware())

	produkCabang.GET("/riwayat", producController.GetAllRiwayat, m.JWTMiddleware())
}
