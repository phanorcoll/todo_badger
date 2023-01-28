package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//Task holds the properties for task instances
type Task struct {
	Id          string `json:"id,omitempty"`
  Title       string `json:"title,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	Completed   bool   `json:"completed,omitempty" validate:"required"`
	Priority    string `json:"priority,omitempty" validate:"required"`
	UserID      string `json:"userId,omitempty" validate:"required"`
	Type        string `json:"type"`
}

var TypeTask = "TASK"

//ListTasks returns the list of tasks
func ListTasks(c echo.Context) error {
	listTasks := []Task{}
	DB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				tTask := Task{}
				_ = json.Unmarshal(val, &tTask)
				if tTask.Type == TypeTask {
					listTasks = append(listTasks, tTask)
				}
				return nil
			})
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
		}
		return nil
	})
	return c.JSON(http.StatusOK, listTasks)
}

//CreateTask creates a new task
func CreateTask(c echo.Context) error {
	task := Task{}
	task.Type = TypeTask
	task.Id = uuid.New().String()[:8]
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
  //validate requird fields are present in task
	if err := validate.Struct(task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	t, err := json.Marshal(task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := DB.Set([]byte(task.Id), []byte(t)); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, task)
}

//GetTask returns data from a particular user
func GetTask(c echo.Context) error {
	taskID := c.Param("taskId")
	v, err := DB.Get([]byte(taskID))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Task not found")
	}

	tTask := Task{}
	_ = json.Unmarshal(v, &tTask)
	return c.JSON(http.StatusOK, tTask)
}

//UpdateTask updates user data base on body sent
func UpdateTask(c echo.Context) error {
	//search for task
	taskID := c.Param("taskId")
	v, err := DB.Get([]byte(taskID))
	if err != nil {
		return c.JSON(http.StatusNotFound, "task not found")
	}
	task := Task{}
	err = json.Unmarshal(v, &task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer c.Request().Body.Close()
	//Get the body fro the request and replace the content of task
	err = json.NewDecoder(c.Request().Body).Decode(&task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}
  //validate requird fields are present in task
	if err := validate.Struct(task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	mTask, err := json.Marshal(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := DB.Set([]byte(task.Id), []byte(mTask)); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, task)
}

//DeleteTask deletes a particular user for the DB
func DeleteTask(c echo.Context) error {
	taskID := c.Param("taskId")
	if err := DB.Delete([]byte(taskID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "task deleted")
}
