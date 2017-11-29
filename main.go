package main

import (
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
	configureStatic(e)
	registerMiddlewares(e)
	registerRoutes(e)

	e.Logger.Fatal(e.Start(serverPort))
}

func configureEcho(e *echo.Echo) {
	e.HideBanner = true
}

func configureStatic(e *echo.Echo) {
	e.File("/", "public/index.html")
	e.Static("/assets", "assets")
}

func registerMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func registerRoutes(e *echo.Echo) {

	// Login
	e.POST(routes.EndpointLogin, routes.Login)

	// Users
	routes.JWTRoute(e, routes.EndpointGetUsers).GET(empty, routes.GetUsers)
	routes.JWTRoute(e, routes.EndpointCreateUser).POST(empty, routes.CreateUser)
	routes.JWTRoute(e, routes.EndpointUpdateUser).PUT(empty, routes.UpdateUser)
	routes.JWTRoute(e, routes.EndpointDeleteUser).DELETE(empty, routes.DeleteUser)
}
