package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/charmbracelet/charm/kv"
	"github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/phanorcoll/todo_badger/api/database"
)

//User holds the properties for user instances
type User struct {
	Id       string `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type UserPublic struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

var DB *kv.KV = database.CreateClient()

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
	listUsers := []UserPublic{}
	DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close() //nolint:errcheck
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			// k := item.Key()
			err := item.Value(func(v []byte) error {
				tuser := UserPublic{}
				_ = json.Unmarshal(v, &tuser)
				// fmt.Printf("%s - %s\n", k, v)
				listUsers = append(listUsers, tuser)
				return nil
			})
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
	return c.JSON(http.StatusOK, listUsers)
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
	user.Id = uuid.New().String()[:8]
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Fatalf("Failed reading the request body %s\n", err)
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	u, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Cannot Marshal User %s\n", err)
	}
	if err := DB.Set([]byte(user.Id), []byte(u)); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, user)
}

//GetUser returns data from a particular user
func GetUser(c echo.Context) error {
	userID := c.Param("userId")
	v, err := DB.Get([]byte(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}

	tuser := User{}
	_ = json.Unmarshal(v, &tuser)

	fmt.Printf("Value form badger is %s\n", v)
	return c.JSON(http.StatusOK, tuser)
}

//UpdateUser updates user data base on body sent
func UpdateUser(c echo.Context) error {
	//search for current user
	userID := c.Param("userId")
	v, err := DB.Get([]byte(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	currentUser := UserPublic{}
	err = json.Unmarshal(v, &currentUser)
	if err != nil {
		log.Fatalf("Cannot Marshal currentUser %s\n", err)
	}

	defer c.Request().Body.Close()

	//Get the body from the request and replace the content of currentUser
	err = json.NewDecoder(c.Request().Body).Decode(&currentUser)
	if err != nil {
		log.Fatalf("Failed reading the request body %s\n", err)
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
	mCurrentUser, err := json.Marshal(currentUser)
	if err != nil {
		log.Fatalf("Cannot Marshal User %s\n", err)
	}

	if err := DB.Set([]byte(currentUser.Id), []byte(mCurrentUser)); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, currentUser)
}

//DeleteUser deletes a particular user for the DB
func DeleteUser(c echo.Context) error {
	userID := c.Param("userId")
	if err := DB.Delete([]byte(userID)); err != nil {
		log.Fatalf("Error deleting a record: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "user deleted")
}
