package web

import (
	"fmt"
	"net/http"
	"waypoint/pkg/models"

	"github.com/labstack/echo/v4"
)

func IndexRoutes(e *echo.Echo) {
	e.GET("/", index)
	e.GET("/users", getUsers)
}

func getUsers(c echo.Context) error {
	users, err := models.GetUsers()
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	return c.Render(http.StatusOK, "users.html", users)
}

func index(c echo.Context) error {
	test := [5]string{
		"hello",
		"there",
		"how",
		"are",
		"you",
	}
	return c.Render(http.StatusOK, "index.html", test)
}
