package controllers

import "github.com/kataras/iris/v12"

func CreateRoutes(irisApp *iris.Application) {
	healthController := NewHealthController()

	irisApp.Get("/health", healthController.Get)

	irisApp.Handle("ALL", "/*", func(ctx iris.Context) {
		ctx.StatusCode(404)
	})
}
