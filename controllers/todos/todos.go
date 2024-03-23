package todos

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"syntaxdata.com/todo_app/components"
	"syntaxdata.com/todo_app/models"
)

func Index(c echo.Context) error {
	var todos []models.Todo
	models.DB.Find(&todos)
	return c.Render(http.StatusOK, "todos/index.html", map[string][]models.Todo{
		"todos": todos,
	})
}

func Create(c echo.Context) error {
	var todo models.Todo
	err := c.Bind(&todo)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	models.DB.Create(&todo)
	// c.Logger().Fatal(todo.Title)
	c.Response().Writer.WriteHeader(http.StatusOK)
	return components.Todo(todo).Render(c.Request().Context(), c.Response().Writer)
}

func Destroy(c echo.Context) error {
	models.DB.Delete(&models.Todo{}, c.Param("id"))
	return c.String(http.StatusOK, "")
}
