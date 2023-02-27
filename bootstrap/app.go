package bootstrap

import (
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/foundation"

	"go-app/config"
)

func Boot() {
	app := foundation.NewApplication()

	//Bootstrap the config.
	config.Boot()

	//Bootstrap the application
	app.Boot()

	facades.Log.Errorf("")
}
