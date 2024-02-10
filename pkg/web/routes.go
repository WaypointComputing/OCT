package web

import (
	"fmt"
	"net/http"
	"time"
	"waypoint/pkg/models"
	"waypoint/pkg/utils"
	"waypoint/pkg/web/mw"

	"github.com/labstack/echo/v4"
)

func IndexRoutes(e *echo.Echo) {
	e.GET("/", index)
	e.GET("/login", loginPage)
	e.POST("/login", login)

	authOnly := e.Group("", mw.Auth)
	authOnly.GET("/users", getUsers)
	authOnly.GET("/cookie", cookieTest)
	authOnly.GET("/get-cookie", getCookie)
	authOnly.POST("/submit-cookie", submitCookie)
}

func getUsers(c echo.Context) error {
	utils.Log("HANDLER - getUsers")

	users, err := models.GetUsers()
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	return c.Render(http.StatusOK, "users.html", users)
}

func index(c echo.Context) error {
	utils.Log("HANDLER - index")

	return c.Render(http.StatusOK, "index.html", nil)
}

func loginPage(c echo.Context) error {
	utils.Log("HANDLER - loginPage")

	return c.Render(http.StatusOK, "login.html", nil)
}

func login(c echo.Context) error {
	utils.Log("HANDLER - login")

	email := c.FormValue("email")
	pwd := c.FormValue("pwd")

	user, err := models.Login(email, pwd)
	if err != nil {
		return err
	}

	if user == nil {
		return c.String(http.StatusUnauthorized, "Incorrect username and password")
	}

	authCookie := new(http.Cookie)
	authCookie.Name = mw.AUTH_COOKIE
	authCookie.Value = "logged in whoohooo"
	authCookie.SameSite = http.SameSiteNoneMode
	authCookie.Secure = true
	authCookie.Expires = time.Now().Add(time.Minute)

	c.SetCookie(authCookie)

	return c.HTML(
		http.StatusOK,
		"<div>Logged in.</div><div><a href='/'>Home</a></div>",
	)
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

	cookie, err := c.Cookie("waypoint-testing")
	if err != nil {
		return c.Render(http.StatusOK, "cookie_test.html", "We could not find your cookie!")
	}

	message := fmt.Sprintf(
		"You have a cookie called %v with the value: %#v",
		cookie.Name,
		cookie.Value,
	)

	return c.Render(http.StatusOK, "cookie_test.html", message)
}
