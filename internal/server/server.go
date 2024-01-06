package server

import (
	"MicroserviceGo/internal/database"
	"MicroserviceGo/internal/models"
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	GetAllCustomers(ctx echo.Context) error
	AddCustomer(ctx echo.Context) error
	GetCustomerById(ctx echo.Context) error

	GetAllProducts(ctx echo.Context) error
	AddProduct(ctx echo.Context) error
	GetProductById(ctx echo.Context) error

	GetAllServices(ctx echo.Context) error
	AddService(ctx echo.Context) error
	GetServiceById(ctx echo.Context) error

	GetAllVendors(ctx echo.Context) error
	AddVendor(ctx echo.Context) error
	GetVendorById(ctx echo.Context) error
}
type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (server *EchoServer) Start() error {
	if err := server.echo.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server shutdown occurred: %s", err)
		return err
	}
	return nil
}

func (server *EchoServer) registerRoutes() {
	server.echo.GET("/readiness", server.Readiness)
	server.echo.GET("/liveness", server.Liveness)

	customerGroup := server.echo.Group("/customers")
	customerGroup.GET("", server.GetAllCustomers)
	customerGroup.POST("", server.AddCustomer)
	customerGroup.GET("/:id", server.GetCustomerById)

	productGroup := server.echo.Group("/products")
	productGroup.GET("", server.GetAllProducts)
	productGroup.POST("", server.AddProduct)
	productGroup.GET("/:productId", server.GetProductById)

	vendorGroup := server.echo.Group("/vendors")
	vendorGroup.GET("", server.GetAllVendors)
	vendorGroup.POST("", server.AddVendor)
	vendorGroup.GET("/:vendorId", server.GetVendorById)

	serviceGroup := server.echo.Group("/services")
	serviceGroup.GET("", server.GetAllServices)
	serviceGroup.POST("", server.AddService)
	serviceGroup.GET("/:servicesId", server.GetServiceById)
}

func (server *EchoServer) Readiness(ctx echo.Context) error {
	ready := server.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "Ok"})
	}
	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

func (server *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "Ok"})
}
