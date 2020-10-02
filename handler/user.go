package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "api get user")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "api update user")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "api delete user")
}
