package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/un-versed/base_api/application"
	"github.com/un-versed/base_api/controllers"
	"github.com/un-versed/base_api/db"
	"github.com/un-versed/base_api/helpers"
)

func initLogger(logLevel logrus.Level) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevel)
}

func main() {
	var serverPort = flag.Int("p", helpers.GetDefaultPort(), "web server http port")
	var pgConnString = flag.String("pg", helpers.GetDefaultDatabaseConnectionString(), "Postgres connection string")
	var logAll = flag.Bool("log-all", false, "Log all messages (trace level)")

	flag.Parse()

	logLevel := logrus.InfoLevel
	if *logAll {
		logLevel = logrus.TraceLevel
	}

	err := db.Open(*pgConnString)
	if err != nil {
		logrus.Fatalf("Unable to connect to database: %s\n", err.Error())
	}
	defer db.Close()

	initLogger(logLevel)

	logrus.Debug("START")

	_app := application.App()

	controllers.CreateRoutes(_app.IrisApp())
	_app.RunServer(*serverPort)

}
