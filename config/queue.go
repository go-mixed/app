package config

import (
	"gopkg.in/go-mixed/framework.v1/facades/config"
)

func queueConfig() {
	config.Add("queue", map[string]any{
		// Default Queue Connection Name
		"default": config.Env("QUEUE_CONNECTION", "sync"),
		// result expire in
		"result_expire": config.Env("QUEUE_RESULT_EXPIRE", "3600s"),

		// Queue Connections
		//
		// Here you may configure the connection information for each server that is used by your application.
		// Drivers: "sync", "redis", "amqp"(rabbitmq)
		"connections": map[string]any{
			"sync": map[string]any{
				"driver": "sync",
			},
			"redis": map[string]any{
				"driver":     "redis",
				"connection": "default",
				"queue":      config.Env("QUEUE_NAME", "default"),
			},
			"rabbitmq": map[string]any{
				"driver":   "amqp",
				"host":     "127.0.0.1",
				"port":     5672,
				"username": "",
				"password": "",
				"vhost":    "/",

				"queue": config.Env("QUEUE_NAME", "default"),
			},
		},
	})
}
