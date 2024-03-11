package routes

import (
	"fmt"
	"net/http"
	"waypoint/pkg/models/user"
	"waypoint/pkg/utils"
	"waypoint/pkg/web/mw"

	"github.com/labstack/echo/v4"
)

func userRoutes(e *echo.Echo) {
	userGroup := e.Group("/user")

	userGroup.GET("/create", createUserPage)
	userGroup.POST("/create", createUser)

	authOnly := userGroup.Group("", mw.Auth)
	authOnly.GET("/users", getUsers)
}

func createUserPage(c echo.Context) error {
	utils.Log("HANDLER - createUserPage")

	return c.Render(http.StatusOK, "createUser.html", nil)
}

func createUser(c echo.Context) error {
	utils.Log("HANDLER - createUser")

	name := c.FormValue("name")
	email := c.FormValue("email")
	pwd := c.FormValue("pwd")
	pwd = user.HashString(pwd)

	newUser, err := user.CreateUser(name, email, pwd)
	if err != nil {
		return err
	}

	return c.HTML(
		http.StatusOK,
		"<p>User '"+newUser.Name+"' created!</p><a href='/'>Home</a>",
	)
}

func getUsers(c echo.Context) error {
	utils.Log("HANDLER - getUsers")

	currentUser := c.Get(user.CURRENT_USER_KEY).(*user.User)
	if currentUser.Privileges < user.PrivilegeLevelAdmin {
		return c.Render(http.StatusUnauthorized, "unauthorized.html", nil)
	}

	users, err := user.GetUsers()
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	return c.Render(http.StatusOK, "users.html", users)
}
