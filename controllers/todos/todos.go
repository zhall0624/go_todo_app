package todos

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"syntaxdata.com/todo_app/components"
	"syntaxdata.com/todo_app/models"
)

func Index(c echo.Context) error {
	var todos []models.Todo
	models.DB.Order("created_at desc").Find(&todos)
	c.Response().Writer.WriteHeader(http.StatusOK)
	return components.TodoIndex(todos).Render(c.Request().Context(), c.Response().Writer)
}

func Create(c echo.Context) error {
	var todo models.Todo
	err := c.Bind(&todo)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	models.DB.Create(&todo)
	c.Response().Writer.WriteHeader(http.StatusOK)
	return components.Todo(todo).Render(c.Request().Context(), c.Response().Writer)
}

func Update(c echo.Context) error {
	var todo models.Todo
	models.DB.First(&todo, c.Param("id"))
	todo.Completed = true
	models.DB.Save(&todo)
	return RenderComponent(c, components.Todo, todo)
}

func Destroy(c echo.Context) error {
	models.DB.Delete(&models.Todo{}, c.Param("id"))
	return c.String(http.StatusOK, "")
}

type componentFunc func(models.Todo) templ.Component

func RenderComponent(c echo.Context, component componentFunc, todo models.Todo) error {
	c.Response().Writer.WriteHeader(http.StatusOK)
	c.Logger().Error(todo)
	return component(todo).Render(c.Request().Context(), c.Response().Writer)
}
