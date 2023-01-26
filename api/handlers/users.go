package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

//User holds the properties for user instances
type User struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

// ListUsers returns the list of users
// ListUsers godoc
//
//	@Summary		List the users in the system
//	@Description	Gets the list of users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	User
//	@Router			/users [get]
func ListUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "list of users")
}

// CreateUser creates a new user
// CreateUser godoc
//
//	@Summary		Creates a new user in the system
//	@Description	Creates a new user 
//	@Tags			users
//	@Accept			json
//	@Produce		json
//  @Param			user	body		User	true	"Create User"
//	@Success		200	{object}	User
//	@Router			/users [post]
func CreateUser(c echo.Context) error {
	user := User{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Fatalf("Failed reading the request body %s\n", err)
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	log.Println(user)
	return c.JSON(http.StatusOK, "user created")
}

//GetUser returns data from a particular user
func GetUser(c echo.Context) error {
	userID := c.Param("userId")
	return c.JSON(http.StatusOK, "getting user ID "+userID)
}

//UpdateUser updates user data base on body sent
func UpdateUser(c echo.Context) error {
	user := User{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Fatalf("Failed reading the request body %s\n", err)
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	return c.JSON(http.StatusOK, user)
}

//DeleteUser deletes a particular user for the DB
func DeleteUser(c echo.Context) error {
	userID := c.Param("userId")
	return c.JSON(http.StatusOK, "Deleting user ID "+userID)
}
