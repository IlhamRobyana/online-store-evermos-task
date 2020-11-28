package product

import (
	"net/http"
	"strconv"

	"github.com/ilhamrobyana/online-store-evermos-task/entity"
	"github.com/ilhamrobyana/online-store-evermos-task/storage"
	"github.com/labstack/echo"
)

var coreInstance *core

func Create(c echo.Context) error {
	r := new(entity.Product)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	productCore := getCore()
	createdProduct, err := productCore.create(*r)

	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusCreated, createdProduct)
}

func GetAll(c echo.Context) error {
	productCore := getCore()
	productList, err := productCore.getAll()
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusOK, productList)
}

func GetByID(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	productCore := getCore()
	product, err := productCore.getByID(productID)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusOK, product)
}

func Update(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	r := new(entity.Product)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	productCore := getCore()
	updatedProduct, err := productCore.update(productID, *r)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusOK, updatedProduct)
}

func Delete(c echo.Context) error {
	productID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	productCore := getCore()
	err := productCore.delete(productID)

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
		productStorage, _ := storage.GetProductStorage(storage.Postgre)

		c.productStorage = productStorage
		coreInstance = c
	}

	return
}
