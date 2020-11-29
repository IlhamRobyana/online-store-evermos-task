package route

import (
	"github.com/ilhamrobyana/online-store-evermos-task/mwcustom"
	"github.com/ilhamrobyana/online-store-evermos-task/order"
	"github.com/labstack/echo"
)

func orderRoute(e *echo.Echo) {
	g := e.Group("/order")
	g.Use(mwcustom.Authorization)
	g.POST("/", order.Create)
	g.GET("/", order.GetAll)
	// g.GET("/:id", order.GetByID)
	g.PUT("/:id", order.Update)
	// g.DELETE("/:id", order.Delete)
}
