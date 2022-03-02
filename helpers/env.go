package helpers

import (
	"os"
	"strconv"
)

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
