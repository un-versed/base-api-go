package db

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var conn *bun.DB

func Conn() *bun.DB {
	return conn
}

func Open(connString string) error {
	config, err := pgx.ParseConfig(connString)
	if err != nil {
		return err
	}

	config.PreferSimpleProtocol = true

	sqldb := stdlib.OpenDB(*config)
	conn = bun.NewDB(sqldb, pgdialect.New())

	return nil
}

func Close() {
	conn.Close()
}
