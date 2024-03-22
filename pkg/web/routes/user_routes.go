package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"waypoint/pkg/auth"
	"waypoint/pkg/models/user"
	"waypoint/pkg/utils"

	"github.com/labstack/echo/v4"
)

func userRoutes(e *echo.Echo) {
	userGroup := e.Group("/user")

	userGroup.GET("/create", createUserPage)
	userGroup.POST("/create", createUser)

	authOnly := userGroup.Group("")
	authOnly.Use(auth.JWTMw())
	authOnly.Use(auth.AuthMw())

	authOnly.GET("/get", getUsers)
	authOnly.GET("/get/:id", getUser)
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
	pwd = auth.HashString(pwd)

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

	currentUser, ok := c.Get(user.CURRENT_USER_KEY).(*user.User)
	if !ok {
		return fmt.Errorf("Error! No user found!")
	}

	if currentUser.Privileges < user.PrivilegeLevelAdmin {
		return c.Render(http.StatusUnauthorized, "unauthorized.html", nil)
	}

	users, err := user.GetUsers()
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	return c.Render(http.StatusOK, "users.html", users)
}

func getUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return err
	}

	u, err := user.GetUser(id)
	if err != nil {
		return err
	}

	log.Println(u)

	return c.String(http.StatusOK, fmt.Sprintf("User: %v, Name: %v, Email: %v, Pwd: %v, Privileges: %v", u.Id, u.Name, u.Email, u.PwdHash, u.Privileges))
}
