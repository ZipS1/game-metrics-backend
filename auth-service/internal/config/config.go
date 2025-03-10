package config

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	SslMode  string `mapstructure:"sslmode"`
	Timezone string `mapstructure:"timezone"`
}

func (c *DatabaseConfig) GetConnectionString() string {
	templateString := "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s"
	return fmt.Sprintf(templateString, c.Host, c.User, c.Password, c.DbName, c.Port, c.SslMode, c.Timezone)
}

type AuthTokensConfig struct {
	JwtExpirationTime          time.Duration `mapstructure:"jwt_expiration_time"`
	RefreshTokenExpirationTime time.Duration `mapstructure:"refresh_token_expiration_time"`
}

type Config struct {
	Port          int              `mapstructure:"port"`
	BaseUriPrefix string           `mapstructure:"base_uri_prefix"`
	AuthTokens    AuthTokensConfig `mapstructure:"auth_tokens"`
	Database      DatabaseConfig   `mapstructure:"database"`
}

func loadConfig(configPath string) (*Config, error) {
	defaults := Config{
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

	viper.SetConfigName(strings.TrimSuffix(filepath.Base(configPath), filepath.Ext(configPath)))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(configPath))
	viper.AutomaticEnv()
	setDefaults(defaults)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := validateConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setDefaults(defaults Config) {
	viper.SetDefault("port", defaults.Port)
	viper.SetDefault("base_uri_prefix", defaults.BaseUriPrefix)
	viper.SetDefault("auth_tokens.jwt_expiration_time", defaults.AuthTokens.JwtExpirationTime)
	viper.SetDefault("auth_tokens.refresh_token_expiration_time", defaults.AuthTokens.RefreshTokenExpirationTime)
	viper.SetDefault("database.port", defaults.Database.Port)
	viper.SetDefault("database.timezone", defaults.Database.Timezone)
}

func validateConfig() error {
	errorMessageTemplate := "no '%s' keyword found in config"

	if !viper.IsSet("database") {
		return fmt.Errorf(errorMessageTemplate, "database")
	}
	requiredDatabaseFields := []string{"host", "user", "password", "dbname", "sslmode"}
	for _, field := range requiredDatabaseFields {
		keywordFullName := fmt.Sprintf("database.%s", field)
		if !viper.IsSet(keywordFullName) {
			return fmt.Errorf(errorMessageTemplate, keywordFullName)
		}
	}

	if !viper.IsSet("auth_tokens") {
		return fmt.Errorf(errorMessageTemplate, "auth_tokens")
	}
	requiredAuthTokensFields := []string{"jwt_expiration_time", "refresh_token_expiration_time"}
	for _, field := range requiredAuthTokensFields {
		keywordFullName := fmt.Sprintf("auth_tokens.%s", field)
		if !viper.IsSet(keywordFullName) {
			return fmt.Errorf(errorMessageTemplate, keywordFullName)
		}
	}

	return nil
}
