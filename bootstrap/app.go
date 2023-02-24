package bootstrap

import (
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/foundation"

	"go-app/config"
)

func Boot() {
	app := &foundation.Application{}

	//Bootstrap the application
	app.Boot()

	//Bootstrap the config.
	config.Boot()

	facades.Log.Errorf("")
}
