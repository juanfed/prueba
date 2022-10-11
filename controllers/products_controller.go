package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"prueba/repositories"
	"prueba/services"
	"strconv"
)

type productController struct {
	service *services.ProductService
}

func ProductMySqlController(mysql *sql.DB) *productController {
	return &productController{
		service: services.MysqlProductService(
			repositories.ProductsRepositories(mysql),
		),
	}
}

func (ctr *productController) GetProduct(c echo.Context) error {
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	product, err := ctr.service.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, product)
}
