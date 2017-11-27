package main

import (
	"net/http"

	"github.com/vsabreu/go-echo-tests/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	serverPort = ":8111"
	empty      = ""
)

func main() {
	e := echo.New()

	configureEcho(e)
	registerMiddlewares(e)
	registerRoutes(e)

	e.Logger.Fatal(e.Start(serverPort))
}

func configureEcho(e *echo.Echo) {
	e.HideBanner = true
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

	// Login
	e.POST(routes.EndpointLogin, routes.Login)

	// Users
	routes.JWTRoute(e, routes.EndpointGetUsers).GET(empty, routes.GetUsers)
	routes.JWTRoute(e, routes.EndpointCreateUser).POST(empty, routes.CreateUser)
	routes.JWTRoute(e, routes.EndpointUpdateUser).PUT(empty, routes.UpdateUser)
	routes.JWTRoute(e, routes.EndpointDeleteUser).DELETE(empty, routes.DeleteUser)
}
