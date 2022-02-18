package app

import (
	"fmt"
	"sync"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"github.com/un-versed/base_api/controllers"
)

var appInstance *application
var appOnce sync.Once

type application struct {
	iris *iris.Application
}

func App() *application {
	appOnce.Do(func() {
		appInstance = &application{iris: iris.New()}
		appInstance.initialize()
	})
	return appInstance
}

func (app *application) initialize() {
	app.iris.Logger().SetLevel("info")
	app.iris.Use(recover.New())
	app.iris.Use(logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		// Columns:            false,
		MessageContextKeys: nil,
		MessageHeaderKeys:  nil,
		LogFunc: func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			l := fmt.Sprintf("ACCESS [%s] %s - %s %s",
				ip,
				status,
				method,
				path)
			logrus.Info(l)
		},
		Skippers: nil,
	}))

	app.createRoutes()
}

func (app *application) RunServer(port int) error {
	return app.iris.Run(iris.Addr(fmt.Sprintf(":%d", port)), iris.WithoutServerError(iris.ErrServerClosed))
}

func (app *application) ShutdownServer() error {
	return app.iris.Shutdown(context.NewContext(app.iris))
}

func (app *application) IrisApp() *iris.Application {
	return app.iris
}

func (app *application) createRoutes() {
	healthController := controllers.NewHealthController()

	app.iris.Get("/health", healthController.Get)

	app.iris.Handle("ALL", "/*", func(ctx iris.Context) {
		ctx.StatusCode(404)
	})
}
