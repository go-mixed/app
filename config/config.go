package config

// Boot Start all init methods of the current folder to bootstrap all config.
func Boot() {
	appConfig()

	// -------------
	// order these configs
	authConfig()
	cacheConfig()
	corsConfig()
	databaseConfig()
	filesystemConfig()
	grpcConfig()
	jwtConfig()
	loggingConfig()
	mailConfig()
	queueConfig()
}
