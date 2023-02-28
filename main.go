package main

import (
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"gopkg.in/go-mixed/framework.v1/facades/log"
	"gopkg.in/go-mixed/framework.v1/facades/route"

	"go-app/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Start http server by facades.Route.
	go func() {
		if err := route.Run(config.GetString("app.host")); err != nil {
			log.Errorf("Route run error: %v", err)
		}
	}()

	select {}
}
