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
	registerMiddlewares(e)
	configureStatic(e)
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
	e.Use(middleware.CSRF())
}

func registerRoutes(e *echo.Echo) {

	// Login
	e.POST(routes.EndpointLogin, routes.Login)

	// Admin Group
	ag := routes.JWTRoute(e, "/admin")
	ag.File("/users", "public/admin/users.html")

	// Users
	ug := routes.JWTRoute(e, "/users")
	ug.GET(empty, routes.GetUsers)
	ug.POST(empty, routes.CreateUser)
	ug.PUT("/:id", routes.UpdateUser)
	ug.DELETE("/:id", routes.DeleteUser)
}
