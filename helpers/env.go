package helpers

import (
	"os"
	"strconv"
)

func GetDefaultDatabaseConnectionString() string {
	res := "postgres://postgres:postgres@localhost:5432/postgres"
	env := os.Getenv("DATABASE_URL")
	if len(env) > 0 {
		res = env
	}

	return res
}

func GetDefaultPort() int {
	res := 8080
	env := os.Getenv("SERVER_PORT")
	if len(env) > 0 {
		p, err := strconv.Atoi(env)
		if err == nil {
			res = p
		}
	}

	return res
}
