package user

import (
	"net/http"

	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	pg "github.com/ilhamrobyana/online-store-evermos-task/pg_storage"
	"github.com/labstack/echo"
)

var coreInstance *core

func Signup(c echo.Context) (e error) {
	r := new(entity.SignupRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	user := entity.User{}
	user.Username = r.Username
	user.Password = r.Password
	userCore := getCore()
	response, e := userCore.signup(user)
	if e != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": e.Error()})
	}
	return c.JSON(http.StatusCreated, response)
}

func Login(c echo.Context) (e error) {
	r := new(entity.LoginRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	authCore := getCore()
	response, err := authCore.login(r.Username, r.Password)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}
	return c.JSON(http.StatusOK, response)
}

func getCore() (c *core) {
	c = coreInstance

	if c == nil {
		c = new(core)
		c.userStore = pg.User{}
		coreInstance = c
	}

	return
}
