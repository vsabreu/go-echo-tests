package routes

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// JWTSecret holds jwt secret in byte array
var JWTSecret = []byte(",787[69.6;4^3%7@8?;&")

// TokenDuration determines for how long a token will live
const TokenDuration = time.Hour * 2

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// JWTRoute returns a new Echo Group with a given endpoint configured
func JWTRoute(e *echo.Echo, endpoint string) *echo.Group {
	r := e.Group(endpoint)

	r.Use(middleware.JWTWithConfig(
		middleware.JWTConfig{
			Claims:      &jwtCustomClaims{},
			SigningKey:  JWTSecret,
			TokenLookup: "cookie:token",
		}))

	return r
}
