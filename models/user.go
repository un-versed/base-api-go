package models

import (
	"context"

	"github.com/un-versed/base_api/db"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int64  `json:"id" bun:",pk,autoincrement"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Serialize() map[string]interface{} {
	s := make(map[string]interface{})

	s["id"] = u.ID
	s["email"] = u.Email

	return s
}

func GetUsers() ([]User, error) {
	db := db.Conn()

	var users []User

	err := db.NewSelect().Model(&users).OrderExpr("id ASC").Scan(context.Background())
	if err != nil {
		return users, err
	}

	return users, err
}

func GetUser(id int64) (User, error) {
	db := db.Conn()

	user := User{}

	err := db.NewSelect().Model(&user).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return user, err
	}

	return user, err
}

func NewUser(u *User) (*User, error) {
	db := db.Conn()

	_, err := db.NewInsert().Model(u).Exec(context.Background())
	if err != nil {
		return u, err
	}

	return u, err
}

func UpdateUser(u *User) (*User, error) {
	db := db.Conn()

	_, err := db.NewUpdate().Model(u).WherePK().Exec(context.Background())

	return u, err
}

func DeleteUser(u *User) error {
	db := db.Conn()

	_, err := db.NewDelete().Model(u).WherePK().Exec(context.Background())

	return err
}
