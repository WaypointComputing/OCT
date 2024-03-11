package mw

import (
	"fmt"
	"net/http"
	"waypoint/pkg/models/user"
	"waypoint/pkg/utils"

	"github.com/labstack/echo/v4"
)

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

		authCookie, err := c.Cookie(user.AUTH_COOKIE)

		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return next(c)
		}

		loggedUser, err := user.GetUserFromSession(authCookie.Value)
		if err != nil {
			return err
		}
		if loggedUser == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return next(c)
		}

		c.Set(user.CURRENT_USER_KEY, loggedUser)

		utils.Log(fmt.Sprintf("Authenticated with %#v", authCookie.Value))

		return next(c)
	}
}
