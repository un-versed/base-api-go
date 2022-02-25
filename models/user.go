package models

import (
	"github.com/un-versed/base_api/db"
)

type User struct {
	XID      int64  `json:"id" db:"id"` // 0
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Serialize() map[string]interface{} {
	s := make(map[string]interface{})

	s["id"] = u.XID
	s["email"] = u.Email

	return s
}

func GetUsers() ([]User, error) {
	var ul []User

	res, err := db.Query("SELECT * FROM users ORDER BY id ASC;")
	if err != nil {
		return ul, err
	}

	err = res.ScanAll(&ul)

	return ul, err
}

func GetUser(id int64) (User, error) {
	u := User{}

	res, err := db.Query("SELECT * FROM users WHERE id = $1 ORDER BY id ASC;", id)
	if err != nil {
		return u, err
	}

	err = res.ScanFirst(&u)
	return u, err
}

func NewUser(user *User) (User, error) {
	u := User{}

	res, err := db.Query("INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *;", user.Email, user.Password)
	if err != nil {
		return u, err
	}

	err = res.ScanFirst(&u)

	return u, err
}

func UpdateUser(u *User) error {
	_, err := db.Query("UPDATE users SET email=$1, password=$2 WHERE id=$3;", u.Email, u.Password, u.XID)
	return err
}

func DeleteUser(u *User) error {
	_, err := db.Query("DELETE FROM users WHERE id=$1;", u.XID)
	return err
}
