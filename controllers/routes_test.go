package controllers

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
	"github.com/un-versed/base_api/application"
)

func TestNotFoundRoutes(t *testing.T) {
	ia := application.App().IrisApp()
	CreateRoutes(ia)

	e := httptest.New(t, ia)

	e.GET("/foo-bar").Expect().Status(httptest.StatusNotFound)
}
