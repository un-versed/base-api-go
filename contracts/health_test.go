package contracts

import (
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/un-versed/base_api/app"
)

func TestGetHealth(t *testing.T) {
	a := app.App()
	i := a.IrisApp()
	e := httptest.New(t, i)

	e.GET("/health").Expect().Status(httptest.StatusOK).JSON().Equal(iris.Map{"status": "ok"})
}
