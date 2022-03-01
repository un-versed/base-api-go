package db

import (
	"flag"
	"reflect"
	"testing"

	"github.com/un-versed/base_api/helpers"
)

var pgConnString = flag.String("pg", helpers.GetDefaultDatabaseConnectionString(), "Postgres connection string")

func TestOpen(t *testing.T) {
	err := Open(*pgConnString)
	if err != nil {
		t.Errorf("Unable to connect to database: %s\n", err.Error())
	}

	err = Open("whatever")
	if err == nil {
		t.Errorf("Connecting with wrong strin")
	}
}

func TestConn(t *testing.T) {
	conn := Conn()
	if reflect.TypeOf(conn).String() != "*sqlx.DB" {
		t.Errorf("Unable to create to connection pointer.")
	}
}

func TestClose(t *testing.T) {
	Close()
}
