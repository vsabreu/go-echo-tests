package routes

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// EndpointLogin holds Login endpoint
const EndpointLogin = "/login"

// Login logs an user in (or not!)
func Login(c echo.Context) error {
	user := c.FormValue("username")
	pass := c.FormValue("passphrase")

	if user == "george" && pass == "harrison" {
		claims := &jwtCustomClaims{
			"George Harrison",
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(TokenDuration).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString(JWTSecret)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{"token": t})
	}

	return echo.ErrUnauthorized
}
