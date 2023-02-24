package routes

import (
	"gopkg.in/go-mixed/framework.v1/contracts/http"
	"gopkg.in/go-mixed/framework.v1/facades"

	"go-app/app/http/controllers"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(200, http.Json{
			"Hello": "Goravel",
		})
	})

	userController := controllers.NewUserController()
	facades.Route.Get("/users/{id}", userController.Show)
}
