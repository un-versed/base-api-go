package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/un-versed/base_api/application"
	"github.com/un-versed/base_api/controllers"
	"github.com/un-versed/base_api/db"
)

func getDefaultPort() int {
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

func getDefaultDatabaseConnectionString() string {
	res := "postgres://postgres:postgres@localhost:5432/postgres"
	env := os.Getenv("DATABASE_URL")
	if len(env) > 0 {
		res = env
	}

	return res
}

func initLogger(logLevel logrus.Level) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevel)
}

func main() {
	var serverPort = flag.Int("p", getDefaultPort(), "web server http port")
	var pgConnString = flag.String("pg", getDefaultDatabaseConnectionString(), "Postgres connection string")
	var logAll = flag.Bool("log-all", false, "Log all messages (trace level)")

	flag.Parse()

	logLevel := logrus.InfoLevel
	if *logAll {
		logLevel = logrus.TraceLevel
	}

	db.Open(*pgConnString)
	defer db.Close()

	initLogger(logLevel)

	logrus.Debug("START")

	_app := application.App()

	controllers.CreateRoutes(_app.IrisApp())
	_app.RunServer(*serverPort)

}
