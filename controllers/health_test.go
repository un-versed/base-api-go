package controllers

import (
	"testing"

	"github.com/kataras/iris/v12/context"
	"github.com/un-versed/base_api/app"
)

func TestHealthController_Get(t *testing.T) {
	a := app.App()
	i := a.IrisApp()
	ctx := context.NewContext(i)

	healthController := NewHealthController()

	healthController.Get(ctx)

	if ctx.GetStatusCode() != 200 {
		t.Errorf("")
	}
}

func TestHealthController_Sum(t *testing.T) {
	t.Run("1 + 1", auxTestHealthController_Sum(1, 1, 2))
	t.Run("2 + 2", auxTestHealthController_Sum(2, 2, 4))
	t.Run("4 + 4", auxTestHealthController_Sum(4, 4, 8))
}

func auxTestHealthController_Sum(a int32, b int32, r int32) func(*testing.T) {
	return func(t *testing.T) {
		s := Sum(a, b)
		if s != r {
			t.Errorf("Expected Sum(a: %d, b: %d) to be equal %d. Got %d", a, b, r, s)
		}
	}
}
