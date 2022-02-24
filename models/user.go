package models

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/un-versed/base_api/db"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int64  `json:"id" bun:",pk,autoincrement"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUsers() ([]User, error) {
	db := db.Open()

	var users []User

	err := db.NewSelect().Model(&users).OrderExpr("id ASC").Scan(context.Background())
	if err != nil {
		logrus.Error("Error selecting users: %v\n", err)
		return users, err
	}

	return users, err
}

func GetUser(id int64) (User, error) {
	db := db.Open()

	user := User{}

	err := db.NewSelect().Model(&user).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		logrus.Error("Error selecting user: %v\n", err)
		return user, err
	}

	return user, err
}

func NewUser(u *User) (*User, error) {
	db := db.Open()

	_, err := db.NewInsert().Model(u).Exec(context.Background())
	if err != nil {
		logrus.Error("Error inserting user: %v\n", err)
		return u, err
	}

	return u, err
}

func UpdateUser(u *User) error {
	db := db.Open()

	_, err := db.NewUpdate().Model(u).WherePK().Exec(context.Background())
	if err != nil {
		logrus.Error("Error updating user: %v\n", err)
	}

	return err
}

func DeleteUser(u *User) error {
	db := db.Open()

	_, err := db.NewDelete().Model(u).WherePK().Exec(context.Background())
	if err != nil {
		logrus.Error("Error deleting user: %v\n", err)
	}

	return err
}
