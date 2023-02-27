package main

import (
	"gopkg.in/go-mixed/framework.v1/facades"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"gopkg.in/go-mixed/framework.v1/facades/log"

	"go-app/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Start http server by facades.Route.
	go func() {
		if err := facades.Route.Run(config.GetString("app.host")); err != nil {
			log.Errorf("Route run error: %v", err)
		}
	}()

	select {}
}
