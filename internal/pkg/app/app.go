package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"umbrella-middleware/internal/app/endpoint"
	"umbrella-middleware/internal/app/mw"
	"umbrella-middleware/internal/app/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()
	a.echo.Use(mw.RoleCheck)
	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Running server...")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
