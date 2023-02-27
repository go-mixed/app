package bootstrap

import (
	"gopkg.in/go-mixed/framework.v1/facades/log"
	"gopkg.in/go-mixed/framework.v1/foundation"

	"go-app/config"
)

func Boot() {
	app := foundation.NewApplication()

	//Bootstrap the config.
	config.Boot()

	//Bootstrap the application
	app.Boot()

	log.Debug("Application termination")
}
