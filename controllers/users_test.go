package controllers

import (
	"flag"
	"strconv"
	"testing"

	"github.com/iris-contrib/httpexpect/v2"
	"github.com/kataras/iris/v12/httptest"
	"github.com/sirupsen/logrus"
	"github.com/un-versed/base_api/application"
	"github.com/un-versed/base_api/db"
	"github.com/un-versed/base_api/helpers"
)

var lastUserId int64
var pgConnString = flag.String("pg", helpers.GetDefaultDatabaseConnectionString(), "Postgres connection string")

func initTest(t *testing.T) *httpexpect.Expect {
	ia := application.App().IrisApp()
	CreateRoutes(ia)

	err := db.Open(*pgConnString)
	if err != nil {
		logrus.Fatalf("Unable to connect to database: %s\n", err.Error())
	}

	e := httptest.New(t, ia)

	return e
}

func TestUsersController_Store(t *testing.T) {
	e := initTest(t)
	defer db.Close()
	defer application.App().ShutdownServer()

	body := map[string]interface{}{
		"email":    "mr_anderson@testing.go",
		"password": "123asd",
	}

	r := e.POST("/users").
		WithJSON(body).
		Expect().
		Status(httptest.StatusOK).
		JSON().
		Object()

	r.Value("email").Equal(body["email"])

	lastUserId = int64(r.Value("id").Number().Raw())
}

func TestUsersController_Index(t *testing.T) {
	e := initTest(t)
	defer db.Close()
	defer application.App().ShutdownServer()

	r := e.GET("/users").
		Expect().
		Status(httptest.StatusOK).
		JSON().
		Array()

	r.NotEmpty()
}

func TestUsersController_Show(t *testing.T) {
	e := initTest(t)
	defer db.Close()
	defer application.App().ShutdownServer()

	id := strconv.FormatInt(lastUserId, 10)
	r := e.GET("/users/" + id).
		Expect().
		Status(httptest.StatusOK).
		JSON().
		Object()

	r.NotEmpty().Value("id")
	r.NotEmpty().Value("email")
}

func TestUsersController_Update(t *testing.T) {
	e := initTest(t)
	defer db.Close()
	defer application.App().ShutdownServer()

	body := map[string]interface{}{
		"email": "neo@testing.go",
	}

	id := strconv.FormatInt(lastUserId, 10)
	r := e.PUT("/users/" + id).
		WithJSON(body).
		Expect().
		Status(httptest.StatusOK).
		JSON().
		Object()

	r.NotEmpty().Value("id").Equal(lastUserId)
	r.NotEmpty().Value("email").Equal(body["email"])
}

func TestUsersController_Delete(t *testing.T) {
	e := initTest(t)
	defer db.Close()
	defer application.App().ShutdownServer()

	id := strconv.FormatInt(lastUserId, 10)
	e.DELETE("/users/" + id).
		Expect().
		Status(httptest.StatusNoContent)
}
