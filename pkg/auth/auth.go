package auth

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"waypoint/pkg/models/user"
	"waypoint/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const (
	SigningKey     string        = "waypoint-oh-yeah"
	AuthCookie     string        = "waypoint-user-token"
	ExpirationTime time.Duration = time.Hour * 2
)

func HashString(input string) string {
	plainText := []byte(input)
	hash := sha512.Sum512(plainText)
	return hex.EncodeToString(hash[:])
}

func JWTMw() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:   []byte(SigningKey),
		TokenLookup:  "cookie:" + AuthCookie,
		ErrorHandler: authErrorHandler,
	})
}

func AuthMw() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			utils.Log("MIDDLEWARE - Auth")

			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return errors.New("Token not found")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return errors.New("Failed to cast *jwt.Token to jwt.MapClaims")
			}

			user_id, err := claims.GetSubject()
			if err != nil {
				return err
			}

			id, err := strconv.ParseInt(user_id, 10, 64)
			if err != nil {
				return err
			}

			u, err := user.GetUser(id)
			if err != nil {
				return err
			}

			c.Set(user.CURRENT_USER_KEY, &u)

			return next(c)
		}
	}
}

func SetJWTCookie(u *user.User, c echo.Context) error {
	token, exp, err := generateToken(u)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = AuthCookie
	cookie.Value = token
	cookie.Expires = exp
	cookie.HttpOnly = true

	c.SetCookie(cookie)

	return nil
}

func generateToken(u *user.User) (string, time.Time, error) {
	expiration := time.Now().Add(ExpirationTime).UTC()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": fmt.Sprint(u.Id),
		"exp": expiration.Unix(),
	})

	signed, err := token.SignedString([]byte(SigningKey))

	return signed, expiration, err
}

func authErrorHandler(c echo.Context, err error) error {
	utils.Log("Failed auth: " + err.Error())
	return c.Redirect(http.StatusSeeOther, c.Echo().Reverse("loginPage"))
}
