package server

import (
	"MicroserviceGo/internal/dberrors"
	"MicroserviceGo/internal/models"
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

func (server *EchoServer) AddService(ctx echo.Context) error {
	service := new(models.Services)
	if err := ctx.Bind(service); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	service, err := server.DB.AddService(ctx.Request().Context(), service)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflicError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, service)
}

func (server *EchoServer) GetServiceById(ctx echo.Context) error {
	Id := ctx.Param("servicesId")
	service, err := server.DB.GetServiceById(ctx.Request().Context(), Id)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, service)
}
