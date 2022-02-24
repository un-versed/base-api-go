package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Open() *pgxpool.Pool {
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://gustavo:123teste@localhost:5432/go_base_api")

	if err != nil {
		logrus.Fatal(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		logrus.Info("Connected to database")
	}

	return dbpool
}

func Close(dbpool *pgxpool.Pool) {
	dbpool.Close()
}
