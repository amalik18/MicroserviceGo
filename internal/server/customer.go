package server

import (
	"MicroserviceGo/internal/dberrors"
	"MicroserviceGo/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (server *EchoServer) GetAllCustomers(ctx echo.Context) error {
	emailAddress := ctx.QueryParam("emailAddress")

	customers, err := server.DB.GetAllCustomers(ctx.Request().Context(), emailAddress)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, customers)
}

func (server *EchoServer) AddCustomer(ctx echo.Context) error {
	customer := new(models.Customer)
	if err := ctx.Bind(customer); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	customer, err := server.DB.AddCustomer(ctx.Request().Context(), customer)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflicError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, customer)
}

func (server *EchoServer) GetCustomerById(ctx echo.Context) error {
	id := ctx.Param("id")
	customer, err := server.DB.GetCustomerById(ctx.Request().Context(), id)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, customer)
}
