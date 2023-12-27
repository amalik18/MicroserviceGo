package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (server *EchoServer) GetAllProducts(ctx echo.Context) error {
	vendorId := ctx.QueryParam("vendorId")
	products, err := server.DB.GetAllProducts(ctx.Request().Context(), vendorId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, products)
}