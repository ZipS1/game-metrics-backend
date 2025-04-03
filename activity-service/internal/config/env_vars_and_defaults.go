package config

import (
	"github.com/spf13/viper"
)

func configureEnvVarsAndDefaults() {
	viper.AutomaticEnv()
	setDefaults(defaultConfig())
}

func defaultConfig() Config {
	return Config{
		Port:              8080,
		PublicUriPrefix:   "/api/activity",
		InternalUriPrefix: "/internal",
		Database: DatabaseConfig{
			Port:     5432,
			Timezone: "UTC",
		},
		AMQP: AMQPConfig{
			Port:    5672,
			Timeout: 5000,
		},
	}
}

func setDefaults(defaults Config) {
	viper.SetDefault("port", defaults.Port)
	viper.SetDefault("public_uri_prefix", defaults.PublicUriPrefix)
	viper.SetDefault("internal_uri_prefix", defaults.InternalUriPrefix)
	viper.SetDefault("database.port", defaults.Database.Port)
	viper.SetDefault("database.timezone", defaults.Database.Timezone)
	viper.SetDefault("amqp.port", defaults.AMQP.Port)
	viper.SetDefault("amqp.timeout", defaults.AMQP.Timeout)
}
