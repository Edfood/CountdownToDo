package main

import (
	"github.com/CountdownToDo/backend/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/assets", "public/assets")

	// e.File("/", "public/index.html")
	// e.File("/signup", "public/signup.html")
	e.POST("/signup", handler.Signup)
	// e.File("/login", "public/signin.html")
	e.POST("/signin", handler.Signin)
	// e.File("/tasks", "public/todos.html")

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config))
	api.POST("/tasks", handler.AddTask)
	api.GET("/tasks", handler.GetTasks)
	api.DELETE("/tasks/:id", handler.DeleteTask)

	return e
}
