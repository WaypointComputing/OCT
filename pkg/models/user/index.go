package user

import (
	"waypoint/pkg/db"
)

type User struct {
	Id         int64
	Name       string
	Email      string
	PwdHash    string
	Privileges int
}

const (
	PrivilegeLevelUser    int = 1
	PrivilegeLevelCreator int = 2
	PrivilegeLevelAdmin   int = 3
)

const (
	ErrorOccurred                int = 0
	IncorrectUsernameAndPassword int = 1
	IncorrectPassword            int = 2
	CorrectUsernameAndPassword   int = 3
)
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

func CreateUser(name string, email string, pwdHash string) (*User, error) {
	result, err := db.Db.Exec(
		"INSERT INTO user (name, email, pwd_hash) VALUES (?, ?, ?)",
		name,
		email,
		pwdHash,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &User{
		Id:         id,
		Name:       name,
		Email:      email,
		PwdHash:    pwdHash,
		Privileges: 1,
	}, nil
}

func GetUsers() (*[]User, error) {
	rows, err := db.QuerySQL("GetUsers")
	if err != nil {
		return nil, err
	}

	users := []User{}

	for rows.Next() {
		var id int64
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

func GetUser(id int64) (User, error) {
	var u User
	row := db.Db.QueryRow("SELECT * FROM user WHERE id = ?", id)
	err := row.Scan(&u.Id, &u.Name, &u.Email, &u.PwdHash, &u.Privileges)
	return u, err
}

func GetUserByEmail(email string) (*User, error) {
	rows, err := db.Db.Query("SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	users := []User{} // Should only be length 0 or 1
	for rows.Next() {
		var id int64
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
