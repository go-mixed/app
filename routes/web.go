package routes

import (
	"gopkg.in/go-mixed/framework.v1/contracts/http"
	"gopkg.in/go-mixed/framework.v1/facades/route"

	"go-app/app/http/controllers"
)

func Web() {
	route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(200, http.Json{
			"Hello": "Goravel",
		})
	})

	userController := controllers.NewUserController()
	route.Get("/users/{id}", userController.Show)
}
