package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type User struct {
	Name string
	Age  int
}

var listUsers = []User{
	{
		Name: "A",
		Age:  1,
	},
	{
		Name: "B",
		Age:  2,
	},
	{
		Name: "C",
		Age:  3,
	},
	{
		Name: "D",
		Age:  4,
	},
	{
		Name: "E",
		Age:  5,
	},
	{
		Name: "F",
		Age:  6,
	},
	{
		Name: "G",
		Age:  7,
	},
	{
		Name: "H",
		Age:  8,
	},
	{
		Name: "I",
		Age:  9,
	},
	{
		Name: "J",
		Age:  10,
	},
}

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "api get user")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "api update user")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "api delete user")
}

func GetAllUser(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	enc := json.NewEncoder(c.Response())

	for _, user := range listUsers {
		if err := enc.Encode(user); err != nil {
			return err
		}

		c.Response().Flush()
		time.Sleep(1 * time.Second)
	}

	return nil
}
