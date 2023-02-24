package main

import (
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/facades/config"

	"go-app/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Start http server by facades.Route.
	go func() {
		if err := facades.Route.Run(config.GetString("app.host")); err != nil {
			facades.Log.Errorf("Route run error: %v", err)
		}
	}()

	select {}
}
