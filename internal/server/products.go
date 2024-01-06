package server

import (
	"MicroserviceGo/internal/dberrors"
	"MicroserviceGo/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (server *EchoServer) GetAllProducts(ctx echo.Context) error {
	vendorId := ctx.QueryParam("vendorId")
	products, err := server.DB.GetAllProducts(ctx.Request().Context(), vendorId)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, products)
}

func (server *EchoServer) AddProduct(ctx echo.Context) error {
	product := new(models.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	product, err := server.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflicError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, product)
}

func (server *EchoServer) GetProductById(ctx echo.Context) error {
	productId := ctx.Param("productId")
	product, err := server.DB.GetProductById(ctx.Request().Context(), productId)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, product)
}
