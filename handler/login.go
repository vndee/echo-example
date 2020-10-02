package handler

import (
	"example.com/m/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func Login(c echo.Context) (err  error) {
	username := c.Get("username").(string)
	admin := c.Get("admin").(bool)

	log.Printf("Login with username &s and admin &v", username, admin)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(2 * time.Minute).Unix()

	t, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		log.Printf("signed token err %v\n", err)
		return err
	}

	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: t,
	})
}
