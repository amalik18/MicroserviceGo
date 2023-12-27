package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (server *EchoServer) GetAllVendors(ctx echo.Context) error {
	vendors, err := server.DB.GetAllVendors(ctx.Request().Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, vendors)
}
