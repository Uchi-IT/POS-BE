package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {

	base := e.Group("v1/api")

	UserRoute(db, base)
	CabangRoute(db, base)
	ProdukCabangRoute(db, base)
	ProdukGudangRoute(db, base)
	TransactionRoute(db, base)
}
