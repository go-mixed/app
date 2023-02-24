package providers

import (
	"gopkg.in/go-mixed/framework.v1/facades"

	"go-app/app/http"
	"go-app/routes"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register() {
	//Add HTTP middlewares
	kernel := http.Kernel{}
	facades.Route.GlobalMiddleware(kernel.Middleware()...)
}

func (receiver *RouteServiceProvider) Boot() {
	//Add routes
	routes.Web()
}
