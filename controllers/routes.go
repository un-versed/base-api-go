package controllers

import "github.com/kataras/iris/v12"

func CreateRoutes(irisApp *iris.Application) {
	healthController := NewHealthController()
	usersController := NewUsersController()

	// GET: http://localhost:{port}/health
	irisApp.Get("/health", healthController.Get)

	usersApi := irisApp.Party("/users")
	{
		// GET: http://localhost:{port}/users
		usersApi.Get("/", usersController.Index)

		// GET: http://localhost:{port}/users/:id
		usersApi.Get("/{id:int64}", usersController.Show)

		// POST: http://localhost:{port}/users
		usersApi.Post("/", usersController.Store)

		// PUT: http://localhost:{port}/users/:id
		usersApi.Put("/{id:int64}", usersController.Update)

		// PUT: http://localhost:{port}/users/:id
		usersApi.Delete("/{id:int64}", usersController.Delete)
	}

	irisApp.Handle("ALL", "/*", func(ctx iris.Context) {
		ctx.StatusCode(404)
	})
}
