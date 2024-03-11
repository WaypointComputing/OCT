package routes

import (
	"net/http"
	"waypoint/pkg/utils"

	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) {
	e.GET("/", index)

	loginRoutes(e)
	cookieRoutes(e)
	userRoutes(e)
}

func index(c echo.Context) error {
	utils.Log("HANDLER - index")

	return c.Render(http.StatusOK, "index.html", nil)
}
