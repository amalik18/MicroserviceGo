package server

import (
	"MicroserviceGo/internal/database"
	"MicroserviceGo/internal/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error
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
	if err := server.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server shutdown occurred: %s", err)
		return err
	}
	return nil
}

func (server *EchoServer) registerRoutes() {
	server.echo.GET("/readiness", server.Readiness)
	server.echo.GET("/liveness", server.Liveness)
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
