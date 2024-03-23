package config

import (
	"github.com/labstack/echo/v4"
	"syntaxdata.com/todo_app/models"
)

func Setup() {
	models.ConnectDatabase()
	e := echo.New()
	Middleware(e)
	Renderer(e)
	Routes(e)
}
