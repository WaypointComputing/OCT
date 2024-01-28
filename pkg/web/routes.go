package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		test := [5]string{
			"hello",
			"there",
			"how",
			"are",
			"you",
		}
		return c.Render(http.StatusOK, "index.html", test)
	})
}
