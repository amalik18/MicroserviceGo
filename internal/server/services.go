package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (server *EchoServer) GetAllServices(ctx echo.Context) error {
	services, err := server.DB.GetAllServices(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, services)
}
