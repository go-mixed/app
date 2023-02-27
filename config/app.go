package config

import (
	"gopkg.in/go-mixed/framework.v1/auth"
	"gopkg.in/go-mixed/framework.v1/cache"
	"gopkg.in/go-mixed/framework.v1/console"
	"gopkg.in/go-mixed/framework.v1/contracts"
	"gopkg.in/go-mixed/framework.v1/database"
	"gopkg.in/go-mixed/framework.v1/event"
	"gopkg.in/go-mixed/framework.v1/facades/config"
	"gopkg.in/go-mixed/framework.v1/filesystem"
	"gopkg.in/go-mixed/framework.v1/grpc"
	"gopkg.in/go-mixed/framework.v1/http"
	"gopkg.in/go-mixed/framework.v1/mail"
	"gopkg.in/go-mixed/framework.v1/queue"
	"gopkg.in/go-mixed/framework.v1/route"
	"gopkg.in/go-mixed/framework.v1/schedule"
	"gopkg.in/go-mixed/framework.v1/validation"

	"go-app/app/providers"
)

func appConfig() {
	config.Add("app", map[string]any{
		// Application Name
		//
		// This value is the name of your application. This value is used when the
		// framework needs to place the application's name in a notification or
		// any other location as required by the application or its packages.
		"name": config.Env("APP_NAME", "go-app"),

		// Application Environment
		//
		// This value determines the "environment" your application is currently
		// running in. This may determine how you prefer to configure various
		// services the application utilizes. Set this in your ".env" file.
		"env": config.Env("APP_ENV", "production"),

		// Application Debug Mode
		"debug": config.Env("APP_DEBUG", false),

		// Application Timezone
		//
		// Here you may specify the default timezone for your application, which
		// will be used by the PHP date and date-time functions. We have gone
		// ahead and set this to a sensible default for you out of the box.
		"timezone": "UTC",

		// Encryption Key
		//
		// 32 character string, otherwise these encrypted strings
		// will not be safe. Please do this before deploying an application!
		"key": config.Env("APP_KEY", ""),

		// Application URL
		"url": config.Env("APP_URL", "http://localhost"),

		// Application host, http server listening address.
		"host": config.Env("APP_HOST", "127.0.0.1:3000"),

		// Autoload service providers
		//
		// The service providers listed here will be automatically loaded on the
		// request to your application. Feel free to add your own services to
		// this array to grant expanded functionality to your applications.
		"providers": []contracts.IServiceProvider{
			&console.ServiceProvider{},
			&database.ServiceProvider{},
			&cache.ServiceProvider{},
			&http.ServiceProvider{},
			&route.ServiceProvider{},
			&schedule.ServiceProvider{},
			&event.ServiceProvider{},
			&queue.ServiceProvider{},
			&grpc.ServiceProvider{},
			&mail.ServiceProvider{},
			&auth.ServiceProvider{},
			&filesystem.ServiceProvider{},
			&validation.ServiceProvider{},
			&providers.AppServiceProvider{},
			&providers.AuthServiceProvider{},
			&providers.RouteServiceProvider{},
			&providers.GrpcServiceProvider{},
			&providers.ConsoleServiceProvider{},
			&providers.QueueServiceProvider{},
			&providers.EventServiceProvider{},
			&providers.ValidationServiceProvider{},
		},
	})
}
