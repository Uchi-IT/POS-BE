package route

import (
	"uchiiParfume/features/cabang/handler"
	"uchiiParfume/features/cabang/repository"
	"uchiiParfume/features/cabang/service"
	m "uchiiParfume/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CabangRoute(db *gorm.DB, e *echo.Group) {
	cabangRepository := repository.NewCabangRepository(db)
	cabangUseCase := service.NewCabangService(cabangRepository)
	cabangController := handler.NewCabangHandler(cabangUseCase)

	cabang := e.Group("/cabang")
	cabang.POST("", cabangController.CreateCabang, m.JWTMiddleware())
	cabang.GET("", cabangController.GetAllCabang, m.JWTMiddleware())
	cabang.GET("/:id", cabangController.GetById, m.JWTMiddleware())
	cabang.PUT("/:id", cabangController.UpdateCabang, m.JWTMiddleware())
	cabang.DELETE("/:id", cabangController.DeleteCabang, m.JWTMiddleware())
}
