package order

import (
	"net/http"
	"strconv"

	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	pg "github.com/ilhamrobyana/online-store-evermos-task/pg_storage"
	"github.com/labstack/echo"
)

var coreInstance *core

func Create(c echo.Context) error {
	userID := c.Get("id").(uint64)
	r := new(entity.OrderCreateRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	orderCore := getCore()
	createdOrder, err := orderCore.create(r.Items, userID)

	if err != nil {
		httpStatus := http.StatusInternalServerError
		if err.Error() == "record not found" {
			c.JSON(http.StatusOK, map[string]interface{}{"message": "Item is unavaliable"})
		}
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusCreated, createdOrder)
}

func GetAll(c echo.Context) error {
	userID := c.Get("id").(uint64)
	orderCore := getCore()
	orderList, err := orderCore.getAll(userID)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusOK, orderList)
}

func GetByID(c echo.Context) error {
	orderID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	orderCore := getCore()
	order, err := orderCore.getByID(orderID)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusOK, order)
}

func Update(c echo.Context) error {
	orderID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	r := new(entity.Order)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	orderCore := getCore()
	updatedOrder, err := orderCore.update(orderID, *r)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusOK, updatedOrder)
}

func Delete(c echo.Context) error {
	orderID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	orderCore := getCore()
	err := orderCore.delete(orderID)

	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func getCore() (c *core) {
	c = coreInstance

	if c == nil {
		c = new(core)

		c.orderStorage = pg.Order{}
		coreInstance = c
	}

	return
}
