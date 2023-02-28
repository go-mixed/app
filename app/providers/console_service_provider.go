package providers

import (
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
	"gopkg.in/go-mixed/framework.v1/facades/schedule"

	"go-app/app/console"
)

type ConsoleServiceProvider struct {
}

func (receiver *ConsoleServiceProvider) Register() {
	kernel := console.Kernel{}
	schedule.Register(kernel.Schedule())
	artisan.Register(kernel.Commands())
}

func (receiver *ConsoleServiceProvider) Boot() {

}
