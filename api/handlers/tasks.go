package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

//Task holds the properties for task instances
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Priority    string `json:"priority"`
}

//ListTasks returns the list of tasks
func ListTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, "list of task per user")
}

//CreateTask creates a new task
func CreateTask(c echo.Context) error {
	task := Task{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	if err != nil {
		log.Fatalf("Failed reading the request body %s\n", err)
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	log.Println(task)
	return c.JSON(http.StatusOK, "task created")
}

//GetTask returns data from a particular user
func GetTask(c echo.Context) error {
	taskID := c.Param("taskId")
	return c.JSON(http.StatusOK, "getting task "+taskID)
}

//UpdateTask updates user data base on body sent
func UpdateTask(c echo.Context) error {
	taskID := c.Param("taskId")
	task := Task{}
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	if err != nil {
		log.Fatalf("Failed reading the request body %s\n", err)
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	log.Println(task)
	return c.JSON(http.StatusOK, "updating task "+taskID)
}

//DeleteTask deletes a particular user for the DB
func DeleteTask(c echo.Context) error {
	taskID := c.Param("taskId")
	return c.JSON(http.StatusOK, "Deleting task "+taskID)
}
