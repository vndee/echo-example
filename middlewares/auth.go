package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
)

func BasicAuth(username string, password string, c echo.Context) (bool, error) {
	log.Printf("username: %s - password: %s", username, password)
	if username == "admin" && password == "password" {
		c.Set("username", username)
		c.Set("admin", true)
		return true, nil
	}

	if username == "vndee" && password == "password" {
		c.Set("username", username)
		c.Set("admin", false)
		return true, nil
	}

	return false, nil
}
