package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

var conn *pgxpool.Pool

func Conn() *pgxpool.Pool {
	return conn
}

func Open(connString string) error {
	c, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return err
	}

	c.Config().MaxConns = 10

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
