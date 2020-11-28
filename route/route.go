package route

import "github.com/labstack/echo"

func Init(e *echo.Echo) {
	userRoute(e)
	productRoute(e)
}
