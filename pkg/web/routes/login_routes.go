package routes

import (
	"net/http"
	"waypoint/pkg/models/user"
	"waypoint/pkg/utils"

	"github.com/labstack/echo/v4"
)

type LoginResponse struct {
	ErrMsg string
	User   *user.User
}

func loginRoutes(e *echo.Echo) {
	e.GET("/login", loginPage)
	e.POST("/login", login)
}

func loginPage(c echo.Context) error {
	utils.Log("HANDLER - loginPage")

	return c.Render(http.StatusOK, "login.html", nil)
}

func login(c echo.Context) error {
	utils.Log("HANDLER - login")

	email := c.FormValue("email")
	pwd := c.FormValue("pwd")
	pwd = user.HashString(pwd)

	loggedUser, err, status := user.Login(email, pwd)

	switch status {
	case user.ErrorOccurred:
		if err != nil {
			return err
		}
		break
	case user.IncorrectUsernameAndPassword:
		return c.String(http.StatusUnauthorized, "Incorrect username and password")
	case user.IncorrectPassword:
		return c.String(http.StatusUnauthorized, "Incorrect password")
	case user.CorrectUsernameAndPassword:
		break
	}

	authCookie, err := user.CreateSessionCookie(loggedUser)
	if err != nil {
		return err
	}

	c.SetCookie(authCookie)

	return c.HTML(
		http.StatusOK,
		"<div>Logged in as "+loggedUser.Name+"</div><div><a href='/'>Home</a></div>",
	)
}
