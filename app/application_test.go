package app

import (
	"testing"

	"github.com/iris-contrib/httpexpect/v2"
	"github.com/kataras/iris/v12/httptest"
)

func TestCreateServer(t *testing.T) {
	// app := App()
	// e := httptest.New(t, app.iris)

	// e.GET("/health").Expect().Status(httptest.StatusOK).JSON().Equal(iris.Map{"status": "ok"})
}

func TestRunServer(t *testing.T) {
	go func() {
		error := App().RunServer(3000)
		if error != nil {
			t.Errorf("Error creating server")
		}
	}()

	defer App().ShutdownServer()
}

func TestNotFoundRoutes(t *testing.T) {
	app := App()
	e := httptest.New(t, app.iris)

	e.GET("/foobar").Expect().Status(httptest.StatusNotFound)
}

func GetRoute(t *testing.T, r string) *httpexpect.Request {
	app := App()
	e := httptest.New(t, app.iris)

	return e.GET(r)
}
