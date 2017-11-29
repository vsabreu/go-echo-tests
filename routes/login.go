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
		exp := time.Now().Add(TokenDuration)
		claims := &jwtCustomClaims{
			"George Harrison",
			true,
			jwt.StandardClaims{
				ExpiresAt: exp.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString(JWTSecret)
		if err != nil {
			return err
		}

		c.SetCookie(createCookie(t, exp))
		return c.JSON(http.StatusOK, echo.Map{"token": t})
	}

	return echo.ErrUnauthorized
}

func createCookie(t string, exp time.Time) *http.Cookie {
	c := new(http.Cookie)
	c.Name = "token"
	c.Value = t
	c.Expires = exp
	c.HttpOnly = true
	return c
}
