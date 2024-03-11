package user

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
	"waypoint/pkg/db"
)

type User struct {
	Id         int
	Name       string
	Email      string
	PwdHash    string
	Privileges int
}

const (
	ErrorOccurred                int = 0
	IncorrectUsernameAndPassword int = 1
	IncorrectPassword            int = 2
	CorrectUsernameAndPassword   int = 3
)
const AUTH_COOKIE string = "waypoint_user_token"
const CURRENT_USER_KEY string = "currentUser"

func Login(email string, pwd string) (*User, error, int) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, err, ErrorOccurred
	}
	if user == nil {
		return nil, nil, IncorrectUsernameAndPassword
	}

	if user.PwdHash != pwd {
		return nil, nil, IncorrectPassword
	}

	return user, nil, CorrectUsernameAndPassword
}

func GetUsers() (*[]User, error) {
	rows, err := db.QuerySQL("GetUsers")
	if err != nil {
		return nil, err
	}

	users := []User{}

	for rows.Next() {
		var id int
		var name string
		var email string
		var pwdHash string
		var privileges int

		err = rows.Scan(&id, &name, &email, &pwdHash, &privileges)
		if err != nil {
			return nil, err
		}

		user := User{
			id,
			name,
			email,
			pwdHash,
			privileges,
		}

		users = append(users, user)
	}

	return &users, nil
}

func GetUserByEmail(email string) (*User, error) {
	rows, err := db.Db.Query("SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	users := []User{} // Should only be length 0 or 1
	for rows.Next() {
		var id int
		var name string
		var userEmail string
		var pwdHash string
		var privileges int

		err := rows.Scan(&id, &name, &userEmail, &pwdHash, &privileges)
		if err != nil {
			return nil, err
		}

		users = append(users, User{
			Id:         id,
			Name:       name,
			Email:      userEmail,
			PwdHash:    pwdHash,
			Privileges: privileges,
		})
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func HashString(input string) string {
	plainText := []byte(input)
	hash := sha512.Sum512(plainText)
	return hex.EncodeToString(hash[:])
}

func CreateSessionCookie(user *User) (*http.Cookie, error) {
	if user == nil {
		return nil, nil
	}

	currentTime := time.Now().UnixMilli()
	hash := fmt.Sprintf("%v%v", currentTime, user.Email)
	hash = HashString(hash)

	expiration := time.Now().Add(time.Minute)

	_, err := db.Db.Exec(
		"INSERT INTO user_session VALUES (?, ?, ?)",
		hash,
		user.Id,
		expiration.UTC().String(),
	)
	if err != nil {
		return nil, err
	}

	sessionCookie := new(http.Cookie)
	sessionCookie.Name = AUTH_COOKIE
	sessionCookie.Value = hash
	sessionCookie.SameSite = http.SameSiteNoneMode
	sessionCookie.Secure = true
	sessionCookie.Expires = expiration

	return sessionCookie, nil
}

func GetUserFromSession(sessionKey string) (*User, error) {
	rows, err := db.Db.Query(
		"SELECT u.* FROM user u, user_session s WHERE u.id = s.user_id AND s.session_key = ?",
		sessionKey,
	)
	if err != nil {
		return nil, err
	}

	users := []User{} // Should only be length 0 or 1
	for rows.Next() {
		var id int
		var name string
		var email string
		var pwdHash string
		var privileges int

		err = rows.Scan(&id, &name, &email, &pwdHash, &privileges)
		if err != nil {
			return nil, err
		}

		users = append(users, User{
			Id:         id,
			Name:       name,
			Email:      email,
			PwdHash:    pwdHash,
			Privileges: privileges,
		})
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}
