package mwcustom

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilhamrobyana/online-store-evermos-task/helper"
	"github.com/labstack/echo"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := getAuthToken(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			return nil
		}

		tokenClaims, err := helper.ParseToken(token)
		if err != nil ||
			!tokenClaims.VerifyExpiresAt(time.Now().Unix(), true) {
			c.NoContent(http.StatusUnauthorized)
			return nil
		}

		err = setAuthData(c, tokenClaims)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
			return nil
		}

		return next(c)
	}
}

func getAuthToken(c echo.Context) (token string, e error) {
	auth := c.Request().Header.Get("Authorization")
	if len(auth) == 0 {
		e = errors.New("Bearer token not provided")
		return
	}

	token = strings.Split(auth, " ")[1]
	return
}

func setAuthData(c echo.Context, claims jwt.MapClaims) error {
	username, ok := claims["username"].(string)
	if !ok {
		return errors.New("username is not available")
	}

	c.Set("username", username)

	strid, ok := claims["id"].(string)
	if !ok {
		return errors.New("id is not available")
	}

	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		return errors.New("convert failed")
	}
	c.Set("id", id)

	return nil
}
