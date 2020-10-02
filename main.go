package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"example.com/m/handler"
	"example.com/m/middlewares"
)

func main() {
	server := echo.New()
	server.Use(middleware.Logger())

	isLogedIn := middleware.JWT([]byte("mysecretkey"))

	server.GET("/", handler.Hello, isLogedIn)
	server.POST("/login", handler.Login, middleware.BasicAuth(middlewares.BasicAuth))
	server.Logger.Fatal(server.Start(":8888"))
}
