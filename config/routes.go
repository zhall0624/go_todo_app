package config

import (
	"github.com/labstack/echo/v4"
	"syntaxdata.com/todo_app/controllers/todos"
)

func Routes(e *echo.Echo) {
	e.GET("/", todos.Index)
	e.POST("/", todos.Create)
	e.PATCH("/:id", todos.Update)
	e.DELETE("/:id", todos.Destroy)
}
