package application

import (
	"fmt"
	"sync"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
)

var appInstance *Application
var appOnce sync.Once

type Application struct {
	iris *iris.Application
}

func App() *Application {
	appOnce.Do(func() {
		appInstance = &Application{iris: iris.New()}
		appInstance.initialize()
	})
	return appInstance
}

func (app *Application) IrisApp() *iris.Application {
	return app.iris
}

func (app *Application) initialize() {
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
		LogFunc: func(_ time.Time, latency time.Duration, status, ip, method, path string, _ interface{}, _ interface{}) {
			l := fmt.Sprintf("ACCESS [%s] %s - %s %s (%s)",
				ip,
				status,
				method,
				path,
				latency,
			)
			logrus.Info(l)
		},
		Skippers: nil,
	}))
}

func (app *Application) RunServer(port int) error {
	return app.iris.Run(iris.Addr(fmt.Sprintf(":%d", port)), iris.WithoutServerError(iris.ErrServerClosed))
}

func (app *Application) ShutdownServer() error {
	return app.iris.Shutdown(context.NewContext(app.iris))
}
