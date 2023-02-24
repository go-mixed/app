package config

import (
	"gopkg.in/go-mixed/framework.v1/facades/config"
)

func init() {
	config.Add("grpc", map[string]any{
		// Grpc Configuration
		//
		// Configure your server host
		"host": config.Env("GRPC_HOST", ""),

		// Configure your client host and interceptors.
		// Interceptors can be the group name of UnaryClientInterceptorGroups in app/grpc/kernel.go.
		"clients": map[string]any{
			//"user": map[string]any{
			//	"host":         config.Env("GRPC_USER_HOST", ""),
			//	"interceptors": []string{},
			//},
		},
	})
}
