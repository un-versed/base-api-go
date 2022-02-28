package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/un-versed/base_api/application"
	"github.com/un-versed/base_api/controllers"
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
	var logAll = flag.Bool("log-all", false, "Log all messages (trace level)")

	flag.Parse()

	logLevel := logrus.InfoLevel
	if *logAll {
		logLevel = logrus.TraceLevel
	}

	initLogger(logLevel)

	logrus.Debug("START")

	_app := application.App()

	controllers.CreateRoutes(_app.IrisApp())
	_app.RunServer(*serverPort)
}
