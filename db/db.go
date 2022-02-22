package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Open() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=gustavo dbname=go_base_api password=123teste sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} else {
		logrus.Info("Connected to database")
	}

	return db
}

func Close(db *sqlx.DB) error {
	return db.Close()
}
