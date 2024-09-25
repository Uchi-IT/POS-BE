package route

import (
	cr "uchiiParfume/features/cabang/repository"
	pcr "uchiiParfume/features/produkCabang/repository"
	"uchiiParfume/features/transaction/handler"
	"uchiiParfume/features/transaction/repository"
	"uchiiParfume/features/transaction/service"
	ur "uchiiParfume/features/users/repository"
	m "uchiiParfume/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRoute(db *gorm.DB, e *echo.Group) {
	produkCabangRepository := pcr.NewProdukRepository(db)
	cabangRepository := cr.NewCabangRepository(db)
	userRepository := ur.NewUserRepository(db, cabangRepository)
	transactionRepository := repository.NewTransactionRepository(db, produkCabangRepository, userRepository)
	transactionUseCase := service.NewTransactionService(transactionRepository)
	transactionController := handler.NewTransactionHandler(transactionUseCase)

	transaction := e.Group("/transaction")
	transaction.POST("", transactionController.CreateTransaction, m.JWTMiddleware())
	transaction.GET("", transactionController.GetAllTransaction, m.JWTMiddleware())
	transaction.GET("/:id", transactionController.GetSpecificTransaction, m.JWTMiddleware())
	transaction.GET("/items/:id", transactionController.GetAllTransactionItemById, m.JWTMiddleware())
}
