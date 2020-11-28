package user

import (
	"net/http"

	"github.com/hipeid/backend/errcode"
	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/ilhamrobyana/online-store-evermos-task/storage"
	"github.com/labstack/echo"
)

var coreInstance *core

func Signup(c echo.Context) (e error) {
	r := new(entity.SignupRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	user := entity.User{0, r.Username, r.Password}
	userCore := getCore()
	response, e := userCore.signup(user)
	if e != nil {
		httpStatus := http.StatusInternalServerError
		if e.Error() == errcode.UserExists {
			httpStatus = http.StatusBadRequest
		}
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
		userStorage, _ := storage.GetUserStorage(storage.Postgre)

		c.userStore = userStorage
		coreInstance = c
	}

	return
}
