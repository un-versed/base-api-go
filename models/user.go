package models

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/un-versed/base_api/db"
)

type User struct {
	ID       int64  `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func GetUsers() ([]User, error) {
	db := db.Open()

	var users []User
	var tmp User

	rows, err := db.Query(context.Background(), "SELECT * FROM users ORDER BY id ASC;")
	if err != nil {
		logrus.Error(err, "Error selecting users")
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&tmp.ID, &tmp.Email, &tmp.Password)
		users = append(users, tmp)
	}

	if rows.Err() != nil {
		logrus.Error("Error reading row: ", err)
		return users, err
	}
	return users, err
}

func GetUser(id int64) (User, error) {
	db := db.Open()

	user := User{}

	row := db.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1 ORDER BY id ASC;", id)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		logrus.Error("Error selecting user:", err)
		return user, err
	}

	return user, err
}

func NewUser(u *User) (User, error) {
	db := db.Open()

	user := User{}
	var id int64

	row := db.QueryRow(context.Background(), `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id;`, &u.Email, &u.Password)
	err := row.Scan(&id)
	if err != nil {
		logrus.Error("Error inserting user:", err)
		return user, err
	}

	user, err = GetUser(id)
	if err != nil {
		logrus.Error("Error fetching user:", err)
		return user, err
	}

	return user, err
}

func UpdateUser(u *User) error {
	db := db.Open()

	_, err := db.Exec(context.Background(), `UPDATE users SET email=$1, password=$2 WHERE id=$3;`, &u.Email, &u.Password, &u.ID)
	if err != nil {
		logrus.Error("Error updating user:", err)
	}

	return err
}

func DeleteUser(u *User) error {
	db := db.Open()

	_, err := db.Exec(context.Background(), `DELETE FROM users WHERE id=$1;`, &u.ID)
	if err != nil {
		logrus.Error("Error deleting user:", err)
	}

	return err
}
