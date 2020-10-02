package main

import (
	"example.com/m/handler"
	"example.com/m/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()
	server.Use(middleware.Logger())

	isLoggedIn := middleware.JWT([]byte("mysecretkey"))

	server.GET("/", handler.Hello, isLoggedIn)
	server.POST("/login", handler.Login, middleware.BasicAuth(middlewares.BasicAuth))
	server.Logger.Fatal(server.Start(":8888"))
}
