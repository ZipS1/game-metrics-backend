package config

import (
	"time"

	"github.com/spf13/viper"
)

func configureEnvVarsAndDefaults() {
	viper.AutomaticEnv()
	viper.BindEnv("auth_tokens.jwt_secret_key", "JWT_SECRET_KEY")
	setDefaults(defaultConfig())
}

func defaultConfig() Config {
	return Config{
		Port:          8080,
		BaseUriPrefix: "/api/auth",
		AuthTokens: AuthTokensConfig{
			JwtExpirationTime:          time.Duration(time.Now().Local().Day()),
			RefreshTokenExpirationTime: time.Duration(time.Now().Year()),
		},
		Database: DatabaseConfig{
			Port:     5432,
			Timezone: "UTC",
		},
	}
}

func setDefaults(defaults Config) {
	viper.SetDefault("port", defaults.Port)
	viper.SetDefault("base_uri_prefix", defaults.BaseUriPrefix)
	viper.SetDefault("auth_tokens.jwt_expiration_time", defaults.AuthTokens.JwtExpirationTime)
	viper.SetDefault("auth_tokens.refresh_token_expiration_time", defaults.AuthTokens.RefreshTokenExpirationTime)
	viper.SetDefault("database.port", defaults.Database.Port)
	viper.SetDefault("database.timezone", defaults.Database.Timezone)
}
