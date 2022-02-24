package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var conn *pgxpool.Pool

func Conn() *pgxpool.Pool {
	return conn
}

func Open(connString string) error {
	c, err := pgxpool.Connect(context.Background(), connString)
	c.Config().MaxConns = 10

	if err != nil {
		logrus.Error(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}

	conn = c
	return nil
}

func Query(sql string, arguments ...interface{}) (QueryResult, error) {
	r, err := conn.Query(context.Background(), sql, arguments...)
	return QueryResult{rows: r}, err
}

func Close() {
	conn.Close()
}
