package controllers

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
	"github.com/un-versed/base_api/app"
)

func TestNotFoundRoutes(t *testing.T) {
	ia := app.App().IrisApp()
	CreateRoutes(ia)

	e := httptest.New(t, ia)

	e.GET("/foo-bar").Expect().Status(httptest.StatusNotFound)
}
