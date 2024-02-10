package models

import (
	"waypoint/pkg/db"
)

type User struct {
	Id         int
	Name       string
	Email      string
	PwdHash    string
	Privileges int
}

func Login(email string, pwd string) (*User, error) {
	users, err := GetUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range *users {
		if user.Email == email && user.PwdHash == pwd {
			return &user, nil
		}
	}

	return nil, nil
}

func GetUsers() (*[]User, error) {
	rows, err := db.Db.QuerySQL("GetUsers")
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
