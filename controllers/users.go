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

	serializedUsers := make([]map[string]interface{}, len(users))

	for i, u := range users {
		serializedUsers[i] = u.Serialize()
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(serializedUsers)
}

func (c *UsersController) Show(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	user, err := models.GetUser(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(user.Serialize())
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
	ctx.JSON(user.Serialize())
}

func (c *UsersController) Update(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	var user models.User
	user.ID = id

	err := ctx.ReadJSON(&user)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	_, err = models.UpdateUser(&user)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(user.Serialize())
}

func (c *UsersController) Delete(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 0)

	var user models.User
	user.ID = id

	err := models.DeleteUser(&user)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(models.NewApiError(err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusNoContent)
}
