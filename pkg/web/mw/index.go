package mw

import (
	"fmt"
	"net/http"
	"waypoint/pkg/utils"

	"github.com/labstack/echo/v4"
)

const AUTH_COOKIE string = "waypoint_user_token"

func TestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		utils.Log("MIDDLEWARE - TestMiddleware")
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		utils.Log("MIDDLEWARE - Auth")

		authCookie, err := c.Cookie(AUTH_COOKIE)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return next(c)
		}

		utils.Log(fmt.Sprintf("Authenticated with %#v", authCookie.Value))

		return next(c)
	}
}
