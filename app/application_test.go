package app

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestCreateServer(t *testing.T) {
	app := App().IrisApp()
	e := httptest.New(t, app)

	e.GET("/foo-bar").Expect().Status(httptest.StatusNotFound)
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
