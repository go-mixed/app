package controllers

import (
	"gopkg.in/go-mixed/framework.v1/contracts/http"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) {
	ctx.Response().Success().Json(200, "", http.Json{
		"Hello": "Laravel",
	})
}
