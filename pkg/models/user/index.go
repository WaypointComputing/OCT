package user

import (
	"database/sql"
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
		user, err := scanUser(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

func GetUser(id int64) (User, error) {
	var u User
	row, err := db.QueryRowSQL("GetUser", id)
	if err != nil {
		return u, err
	}
	return u, row.Scan(&u.Id, &u.Name, &u.Email, &u.PwdHash, &u.Privileges)
}

func GetUserByEmail(email string) (*User, error) {
	rows, err := db.QuerySQL("GetUserByEmail", email)
	if err != nil {
		return nil, err
	}

	users := []User{} // Should only be length 0 or 1
	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func scanUser(rows *sql.Rows) (User, error) {
	var user User

	err := rows.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.PwdHash,
		&user.Privileges,
	)

	return user, err
}
