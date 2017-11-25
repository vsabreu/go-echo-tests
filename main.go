package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const serverPort = 8111

type (
	// Role defines access level
	Role int

	// User represents an application user
	User struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Role   Role   `json:"role"`
		Active bool   `json:"active"`
	}

	// Product represents a product for sell
	Product struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"desc"`
		Price       float64 `json:"price"`
	}
)

var (
	users    map[int]*User
	products map[int]*Product

	usersSeq    int
	productsSeq int
)

const (
	appUser Role = iota + 1
	appAdmin
)

func init() {
	users = make(map[int]*User)
	products = make(map[int]*Product)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Here we sell 5 products. Wanna register to find out?")
	})

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, users)
	})
	e.POST("/users", func(c echo.Context) error {
		u := &User{ID: usersSeq, Role: appUser, Active: false}
		if err := c.Bind(u); err != nil {
			return err
		}

		users[u.ID] = u
		usersSeq++

		return c.JSON(http.StatusCreated, u)
	})
	e.PUT("/users/:id", func(c echo.Context) error {
		nu := &User{}
		if err := c.Bind(nu); err != nil {
			return err
		}

		id, _ := strconv.Atoi(c.Param("id"))
		users[id].Name = nu.Name
		users[id].Email = nu.Email

		return c.JSON(http.StatusOK, users[id])
	})
	e.DELETE("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		delete(users, id)
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", serverPort)))
}
