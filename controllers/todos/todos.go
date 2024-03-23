package todos

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"syntaxdata.com/todo_app/models"
)

func Index(c echo.Context) error {
	var todos []models.Todo
	models.DB.Find(&todos)
	return c.Render(200, "todos/index.html", map[string]interface{}{
		"todos": todos,
	})
}

func Create(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")

	models.DB.Create(&models.Todo{Title: title, Description: description})
	return c.Redirect(http.StatusFound, "/")
}

func Destroy(c echo.Context) error {
	models.DB.Delete(&models.Todo{}, c.Param("id"))
	return c.Redirect(http.StatusFound, "/")
}
