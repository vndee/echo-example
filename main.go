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
	isAdmin := middlewares.IsAdminMiddleware

	server.GET("/", handler.Hello, isLoggedIn)
	server.POST("/login", handler.Login, middleware.BasicAuth(middlewares.BasicAuth))
	server.GET("/admin", handler.Hello, isLoggedIn, isAdmin)

	groupv2 := server.Group("/v2", isLoggedIn)
	groupv2.GET("/hello", handler.Hello2)

	groupUser := server.Group("/api/user")
	groupUser.GET("/get", handler.GetUser)
	groupUser.GET("/update", handler.UpdateUser, isAdmin)
	groupUser.GET("/delete", handler.DeleteUser, isAdmin)
	groupUser.GET("/get_all", handler.GetAllUser)

	server.Logger.Fatal(server.Start(":8888"))
}
