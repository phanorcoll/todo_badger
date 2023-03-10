package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phanorcoll/todo_badger/api/handlers"
	"github.com/phanorcoll/todo_badger/config"
	_ "github.com/phanorcoll/todo_badger/docs"
	"github.com/swaggo/echo-swagger"
)

// @title Todo Badger (Name will Change)
// @version 1.0
// @description Todo application implementing CharmKV(Badger) as database

// @contact.name Phanor Coll
// @contact.url https://www.phanorcoll.com
// @contact.email phanorcoll@gmail.com

// @host localhost:8000
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/login", handlers.Login)

	v1 := e.Group("/api/v1")
	v1.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.EnvVariables.SECRET_KEY),
	}))

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
	gTask.GET("/user/:userId", handlers.ListTaskPerUser)
	gTask.GET("/:taskId", handlers.GetTask)
	gTask.PUT("/:taskId", handlers.UpdateTask)
	gTask.DELETE("/:taskId", handlers.DeleteTask)

	e.Logger.Fatal(e.Start(":" + config.EnvVariables.PORT))
}
