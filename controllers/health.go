package controllers

import (
	"github.com/kataras/iris/v12"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) Get(ctx iris.Context) {
	ctx.StatusCode(200)
	ctx.JSON(iris.Map{"status": "ok"})
}

func Sum(a int32, b int32) int32 {
	return a + b
}
