package models

import (
	"flag"
	"testing"

	"github.com/un-versed/base_api/db"
	"github.com/un-versed/base_api/helpers"
)

var user User = User{XID: 1, Email: "baseapi@test.com", Password: "123"}
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
	case u.XID <= 0:
		t.Errorf("Expected u.XID to be greater than 0")
	}

	user = u
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

	u, err := GetUser(user.XID)
	if err != nil {
		t.Errorf("Unable to get user: %s\n", err.Error())
	}

	switch {
	case u.Email != user.Email:
		t.Errorf("Expected u.Email to be equal %+s. Got %+s", user.Email, u.Email)
	case u.Password != user.Password:
		t.Errorf("Expected u.Password to be equal %+s. Got %+s", user.Password, u.Password)
	case u.XID == 0:
		t.Errorf("Expected u.XID to be greater than 0")
	}
}

func TestUpdateUser(t *testing.T) {
	err := db.Open(*pgConnString)
	if err != nil {
		t.Errorf("Unable to connect to database: %s\n", err.Error())
	}
	defer db.Close()

	ID := user.XID
	newUser := User{XID: ID, Email: "baseapi_update@test.com", Password: "123"}

	u, err := UpdateUser(&newUser)
	if err != nil {
		t.Errorf("Unable to get users: %s\n", err.Error())
	}

	switch {
	case u.Email != newUser.Email:
		t.Errorf("Expected u.Email to be equal %+s. Got %+s", newUser.Email, u.Email)
	case u.Password != newUser.Password:
		t.Errorf("Expected u.Password to be equal %+s. Got %+s", newUser.Password, u.Password)
	case u.XID == 0:
		t.Errorf("Expected u.XID to be greater than 0")
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

	_, err = GetUser(user.XID)
	if err == nil {
		t.Errorf("Unable to delete user: %s\n", err.Error())
	}
}

func TestUserSerialize(t *testing.T) {
	user := User{XID: 1, Email: "test@testing.go", Password: "123asd"}
	serialized := user.Serialize()

	if serialized["password"] != nil {
		t.Errorf("Failed removing password with Serialize().")
	}
}
