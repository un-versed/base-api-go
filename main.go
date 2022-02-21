package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"github.com/un-versed/base_api/app"
	"github.com/un-versed/base_api/controllers"
)

func getDefaultPort() int {
	defaultEnvPort := 8080
	envPort := os.Getenv("SERVER_PORT")
	if len(envPort) > 0 {
		p, err := strconv.Atoi(envPort)
		if err == nil {
			defaultEnvPort = p
		}
	}

	return defaultEnvPort
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
	var logAll = flag.Bool("log-all", false, "Log all messages (trace level)")

	flag.Parse()

	logLevel := logrus.InfoLevel
	if *logAll {
		logLevel = logrus.TraceLevel
	}

	initLogger(logLevel)

	logrus.Debug("START")

	_app := app.App()
	createRoutes(_app)
	_app.RunServer(*serverPort)
}

func createRoutes(a *app.Application) {
	cl := make([]controllers.Router, 0)
	cl = append(cl, controllers.NewHealthController())

	for _, c := range cl {
		c.Route(a.IrisApp())
	}

	a.IrisApp().Handle("ALL", "/*", func(ctx iris.Context) {
		ctx.StatusCode(404)
	})

}
