package route

import (
	"github.com/ilhamrobyana/online-store-evermos-task/mwcustom"
	"github.com/ilhamrobyana/online-store-evermos-task/product"
	"github.com/labstack/echo"
)

func productRoute(e *echo.Echo) {
	g := e.Group("/product")
	g.Use(mwcustom.Authorization)
	g.POST("/", product.Create)
	g.GET("/", product.GetAll)
	g.GET("/:id", product.GetByID)
	g.PUT("/:id", product.Update)
	g.DELETE("/:id", product.Delete)
}
