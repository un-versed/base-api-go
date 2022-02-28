package models

import (
	"flag"
	"testing"

	"github.com/un-versed/base_api/db"
	"github.com/un-versed/base_api/helpers"
)

var user User = User{ID: 1, Email: "baseapi@test.com", Password: "123"}
var pgConnString = flag.String("pg", helpers.GetDefaultDatabaseConnectionString(), "Postgres connection string")

func TestNewUser(t *testing.T) {
	err := db.Open(*pgConnString)
	if err != nil {
		t.Errorf("Unable to connect to database: %s\n", err.Error())
	}
	defer db.Close()

	u, err := NewUser(&user)
	if err != nil {
		t.Errorf("Unable to create user: %s\n", err.Error())
	}

	switch {
	case u.Email != user.Email:
		t.Errorf("Expected u.Email to be equal %+s. Got %+s", user.Email, u.Email)
	case u.Password != user.Password:
		t.Errorf("Expected u.Password to be equal %+s. Got %+s", user.Password, u.Password)
	case u.ID == 0:
		t.Errorf("Expected u.ID to be greater than 0")
	}
}

func TestGetUsers(t *testing.T) {
	err := db.Open(*pgConnString)
	if err != nil {
		t.Errorf("Unable to connect to database: %s\n", err.Error())
	}
	defer db.Close()

	users, err := GetUsers()
	if err != nil {
		t.Errorf("Unable to get users: %s\n", err.Error())
	}

	if len(users) <= 0 {
		t.Errorf("Expected users array length to be greater than 0")
	}
}

func TestGetUser(t *testing.T) {
	err := db.Open(*pgConnString)
	if err != nil {
		t.Errorf("Unable to connect to database: %s\n", err.Error())
	}
	defer db.Close()

	u, err := GetUser(1)
	if err != nil {
		t.Errorf("Unable to get users: %s\n", err.Error())
	}

	switch {
	case u.Email != user.Email:
		t.Errorf("Expected u.Email to be equal %+s. Got %+s", user.Email, u.Email)
	case u.Password != user.Password:
		t.Errorf("Expected u.Password to be equal %+s. Got %+s", user.Password, u.Password)
	case u.ID == 0:
		t.Errorf("Expected u.ID to be greater than 0")
	}
}

func TestUpdateUser(t *testing.T) {
	err := db.Open(*pgConnString)
	if err != nil {
		t.Errorf("Unable to connect to database: %s\n", err.Error())
	}
	defer db.Close()

	newUser := User{ID: 1, Email: "baseapi_update@test.com", Password: "123"}

	u, err := UpdateUser(&newUser)
	if err != nil {
		t.Errorf("Unable to get users: %s\n", err.Error())
	}

	switch {
	case u.Email != newUser.Email:
		t.Errorf("Expected u.Email to be equal %+s. Got %+s", newUser.Email, u.Email)
	case u.Password != newUser.Password:
		t.Errorf("Expected u.Password to be equal %+s. Got %+s", newUser.Password, u.Password)
	case u.ID == 0:
		t.Errorf("Expected u.ID to be greater than 0")
	}
}

func TestDeleteUser(t *testing.T) {
	err := db.Open(*pgConnString)
	if err != nil {
		t.Errorf("Unable to connect to database: %s\n", err.Error())
	}
	defer db.Close()

	err = DeleteUser(&user)
	if err != nil {
		t.Errorf("Unable to delete user: %s\n", err.Error())
	}

	_, err = GetUser(1)
	if err == nil {
		t.Errorf("Unable to delete user: %s\n", err.Error())
	}
}
