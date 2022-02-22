package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/un-versed/base_api/models"
)

type UsersController struct {
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (c *UsersController) Index(ctx iris.Context) {
	users, err := models.GetUsers()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(users)
}

func (c *UsersController) Show(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	users, err := models.GetUser(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(users)
}

func (c *UsersController) Store(ctx iris.Context) {
	var u models.User
	err := ctx.ReadJSON(&u)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	user, err := models.NewUser(&u)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(user)
}

func (c *UsersController) Update(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	var u models.User
	u.ID = id

	err := ctx.ReadJSON(&u)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	err = models.UpdateUser(&u)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(u)
}

func (c *UsersController) Delete(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	var u models.User
	u.ID = id

	err := ctx.ReadJSON(&u)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	err = models.DeleteUser(&u)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusNoContent)
}
