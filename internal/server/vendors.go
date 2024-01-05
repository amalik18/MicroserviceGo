package server

import (
	"MicroserviceGo/internal/dberrors"
	"MicroserviceGo/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (server *EchoServer) GetAllVendors(ctx echo.Context) error {
	vendors, err := server.DB.GetAllVendors(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, vendors)
}

func (server *EchoServer) AddVendor(ctx echo.Context) error {
	vendor := new(models.Vendors)
	if err := ctx.Bind(vendor); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	vendor, err := server.DB.AddVendor(ctx.Request().Context(), vendor)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflicError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, vendor)
}
