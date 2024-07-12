package route

import (
	"uchiiParfume/features/users/handler"
	"uchiiParfume/features/users/repository"
	"uchiiParfume/features/users/service"
	m "uchiiParfume/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute(db *gorm.DB, e *echo.Group) {
	userRepository := repository.NewUserRepository(db)
	userUseCase := service.NewUserService(userRepository)
	userController := handler.NewUserHandler(userUseCase)

	user := e.Group("/user")
	user.POST("", userController.CreateUser)
	user.POST("/login", userController.Login)
	user.GET("", userController.GetAllUser, m.JWTMiddleware())
	user.GET("/profile", userController.GetSpecificUser, m.JWTMiddleware())
	user.DELETE("/:id", userController.DeleteUser, m.JWTMiddleware())
}
