package controllers

import "github.com/kataras/iris/v12"

type Router interface {
	Route(irisApp *iris.Application)
}
