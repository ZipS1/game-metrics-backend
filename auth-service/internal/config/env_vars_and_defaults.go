package config

import (
	"time"

	"github.com/spf13/viper"
)

func configureEnvVarsAndDefaults() error {
	viper.AutomaticEnv()
	if err := viper.BindEnv("jwt_token.public_key_filepath", "JWT_PUBLIC_KEY_PEM_FILEPATH"); err != nil {
		return err
	}
	if err := viper.BindEnv("jwt_token.private_key_filepath", "JWT_PRIVATE_KEY_PEM_FILEPATH"); err != nil {
		return err
	}
	setDefaults(defaultConfig())
	return nil
}

func defaultConfig() Config {
	return Config{
		Port:              8080,
		PublicUriPrefix:   "/api/auth",
		InternalUriPrefix: "/internal",
		JwtToken: JwtTokenConfig{
			JwtExpirationTime:          time.Duration(time.Now().Local().Day()),
			RefreshTokenExpirationTime: time.Duration(time.Now().Year()),
		},
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
	viper.SetDefault("jwt_token.jwt_expiration_time", defaults.JwtToken.JwtExpirationTime)
	viper.SetDefault("jwt_token.refresh_token_expiration_time", defaults.JwtToken.RefreshTokenExpirationTime)
	viper.SetDefault("database.port", defaults.Database.Port)
	viper.SetDefault("database.timezone", defaults.Database.Timezone)
	viper.SetDefault("amqp.port", defaults.AMQP.Port)
	viper.SetDefault("amqp.timeout", defaults.AMQP.Timeout)
}
