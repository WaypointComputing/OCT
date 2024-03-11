package routes

import (
	"fmt"
	"net/http"
	"time"
	"waypoint/pkg/models/user"
	"waypoint/pkg/utils"
	"waypoint/pkg/web/mw"

	"github.com/labstack/echo/v4"
)

func cookieRoutes(e *echo.Echo) {
	authOnly := e.Group("", mw.Auth)
	authOnly.GET("/users", getUsers)
	authOnly.GET("/cookie", cookieTest)
	authOnly.GET("/get-cookie", getCookie)
	authOnly.POST("/submit-cookie", submitCookie)
}

func getUsers(c echo.Context) error {
	utils.Log("HANDLER - getUsers")

	users, err := user.GetUsers()
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	return c.Render(http.StatusOK, "users.html", users)
}

func getCookie(c echo.Context) error {
	utils.Log("HANDLER - getCookie")

	return c.Render(http.StatusOK, "get_cookie.html", nil)
}

func submitCookie(c echo.Context) error {
	utils.Log("HANDLER - submitCookie")

	value := c.FormValue("cookie-val")

	cookie := new(http.Cookie)
	cookie.Name = "waypoint-testing"
	cookie.Value = value
	cookie.Expires = time.Now().Add(time.Minute)
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true

	c.SetCookie(cookie)

	return c.HTML(http.StatusOK, "<h4>You have gained a cookie. It will expire in one minute.</h4>")
}

func cookieTest(c echo.Context) error {
	utils.Log("HANDLER - cookieTest")

	var currentUser *user.User
	currentUser = c.Get(user.CURRENT_USER_KEY).(*user.User)

	cookie, err := c.Cookie("waypoint-testing")
	if err != nil {
		return c.Render(http.StatusOK, "cookie_test.html", "Could not find cookies for "+currentUser.Name+"!")
	}

	message := fmt.Sprintf(
		currentUser.Name+" has a cookie called %v with the value: %#v",
		cookie.Name,
		cookie.Value,
	)

	return c.Render(http.StatusOK, "cookie_test.html", message)
}
