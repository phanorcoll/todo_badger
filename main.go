package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phanorcoll/todo_badger/api/handlers"
	"github.com/phanorcoll/todo_badger/config"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
    return c.JSON(http.StatusOK,"home endpoint")
	})

  v1:=e.Group("/api/v1")

	//User related routes
	gUser := v1.Group("/users")
	gUser.GET("", handlers.ListUsers)
	gUser.POST("", handlers.CreateUser)
	gUser.GET("/:userId", handlers.GetUser)
	gUser.PUT("/:userId", handlers.UpdateUser)
	gUser.DELETE("/:userId", handlers.DeleteUser)

	//Tasks related routes
	gTask := v1.Group("/tasks")
	gTask.GET("", handlers.ListTasks)
  gTask.POST("", handlers.CreateTask)
	gTask.GET("/:taskId", handlers.GetTask)
	gTask.PUT("/:taskId", handlers.UpdateTask)
	gTask.DELETE("/:taskId", handlers.DeleteTask)

	e.Logger.Fatal(e.Start(":" + config.EnvVariables.PORT))
}
