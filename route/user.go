package route

import (
	"github.com/labstack/echo"

	"github.com/ilhamrobyana/online-store-evermos-task/user"
)

func userRoute(e *echo.Echo) {
	g := e.Group("/user")
	g.POST("/signup", user.Signup)
	g.POST("/login", user.Login)
}
