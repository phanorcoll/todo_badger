package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/phanorcoll/todo_badger/api/helper"
	"github.com/phanorcoll/todo_badger/config"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	//search to see if the email exists
	v, err := DB.Get([]byte(email))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"msg": "email not found",
		})
	}

	tuser := User{}
	_ = json.Unmarshal(v, &tuser)

	//verify the password is the same
	check := helper.VerifyPassword(tuser.Password, password)
	if check != true {
		return echo.ErrUnauthorized
	}

	//set custom claims
	claims := &jwtCustomClaims{
		tuser.Name,
		tuser.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	//Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Generate encoded token and send it as response
	t, err := token.SignedString([]byte(config.EnvVariables.SECRET_KEY))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
