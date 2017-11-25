package main

import (
	"fmt"
	"net/http"

	"github.com/vsabreu/go-echo-tests/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const serverPort = 8111

func main() {
	e := echo.New()

	registerMiddlewares(e)
	registerRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", serverPort)))
}

func registerMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func registerRoutes(e *echo.Echo) {
	// Index
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Here we sell 5 products. Wanna register to find out?")
	})

	// Users
	e.GET(routes.EndpointGetUsers, routes.GetUsers)
	e.POST(routes.EndpointCreateUser, routes.CreateUser)
	e.PUT(routes.EndpointUpdateUser, routes.UpdateUser)
	e.DELETE(routes.EndpointDeleteUser, routes.DeleteUser)
}
