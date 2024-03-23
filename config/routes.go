package config

import (
	"github.com/labstack/echo/v4"
	"syntaxdata.com/todo_app/controllers/todos"
)

func Routes(e *echo.Echo) {
	e.GET("/", todos.Index)
	e.POST("/", todos.Create)
	e.POST("/:id/destroy", todos.Destroy)
}
