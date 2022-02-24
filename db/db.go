package db

import (
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func Open() *bun.DB {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		logrus.Fatal(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	} else {
		logrus.Info("Connected to database")
	}
	config.PreferSimpleProtocol = true

	sqldb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}

func Close(db *bun.DB) {
	db.Close()
}
