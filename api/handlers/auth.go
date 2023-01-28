package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
  "github.com/phanorcoll/todo_badger/config"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	//Throws unauthorized error
	if username != "jon" || password != "snow" {
		return echo.ErrUnauthorized
	}

	//set custom claims
	claims := &jwtCustomClaims{
		"John Snow",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	//Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Generate ecndoded token and send it as response
	t, err := token.SignedString([]byte(config.EnvVariables.SECRET_KEY))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
