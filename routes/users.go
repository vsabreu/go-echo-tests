package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/vsabreu/go-echo-tests/models"
)

var (
	users    map[int]*models.User
	usersSeq int
)

func init() {
	users = make(map[int]*models.User)
}

// GetUsers retrieves all users
func GetUsers(c echo.Context) error {
	u := []*models.User{}
	for _, v := range users {
		u = append(u, v)
	}

	return c.JSON(http.StatusOK, u)
}

// CreateUser creates a new user
func CreateUser(c echo.Context) error {
	u := &models.User{ID: usersSeq, Role: models.AppUser, Active: false}
	if err := c.Bind(u); err != nil {
		return err
	}

	users[u.ID] = u
	usersSeq++

	return c.JSON(http.StatusCreated, u)
}

// UpdateUser updates a given user
func UpdateUser(c echo.Context) error {
	nu := &models.User{}
	if err := c.Bind(nu); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = nu.Name
	users[id].Email = nu.Email

	return c.JSON(http.StatusOK, users[id])
}

// DeleteUser deletes a given user
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
