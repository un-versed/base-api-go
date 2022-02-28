package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var conn *sqlx.DB

func Conn() *sqlx.DB {
	return conn
}

func Open(connString string) error {
	c, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return err
	}

	conn = c
	return nil
}

func Close() {
	conn.Close()
}
