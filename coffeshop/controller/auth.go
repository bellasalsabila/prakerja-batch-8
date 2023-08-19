package controller

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const jwtSecret = "e4hBrbL4ZmAmHDEruYeT" // Ganti dengan secret key yang aman

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(jwtSecret),
	})
}

func ExtractToken(authHeader string) string {
	return strings.Replace(authHeader, "Bearer ", "", 1)
}

func Authenticate(username, password string) bool {
	// Implement logic to authenticate user
	// Return true if authentication is successful
	return true
}

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token akan berlaku selama 24 jam

	return token.SignedString([]byte(jwtSecret))
}
