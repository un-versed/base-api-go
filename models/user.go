package models

import (
	"github.com/un-versed/base_api/db"
)

type User struct {
	ID       int64  `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (u *User) Serialize() map[string]interface{} {
	s := make(map[string]interface{})

	s["id"] = u.ID
	s["email"] = u.Email

	return s
}

func GetUsers() ([]User, error) {
	db := db.Conn()
	users := []User{}

	err := db.Select(&users, "SELECT * FROM users ORDER BY id ASC;")

	return users, err
}

func GetUser(id int64) (User, error) {
	db := db.Conn()

	user := User{}

	err := db.Get(&user, "SELECT * FROM users WHERE id = $1 ORDER BY id ASC;", id)

	return user, err
}

func NewUser(u *User) (User, error) {
	db := db.Conn()

	user := User{}

	stmt, err := db.PrepareNamed(`INSERT INTO users (email, password) VALUES (:email, :password) RETURNING id;`)
	if err != nil {
		return user, err
	}

	var id int64

	err = stmt.Get(&id, u)
	if err != nil {
		return user, err
	}

	err = db.Get(&user, "SELECT * FROM users WHERE id = $1 ORDER BY id ASC;", id)
	if err != nil {
		return user, err
	}

	return user, err
}

func UpdateUser(u *User) (User, error) {
	db := db.Conn()

	_, err := db.NamedExec(`UPDATE users SET email=:email, password=:password WHERE id=:id;`, u)
	if err != nil {
		return User{}, err
	}
	user, err := GetUser(u.ID)

	return user, err
}

func DeleteUser(u *User) error {
	db := db.Conn()

	_, err := db.NamedExec(`DELETE FROM users WHERE id=:id;`, u)

	return err
}
