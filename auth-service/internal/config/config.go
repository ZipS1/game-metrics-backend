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
	JwtSecretKey               string        `mapstructure:"jwt_secret_key" env:"JWT_SECRET_KEY"`
	JwtExpirationTime          time.Duration `mapstructure:"jwt_expiration_time"`
	RefreshTokenExpirationTime time.Duration `mapstructure:"refresh_token_expiration_time"`
}

type Config struct {
	DomainName    string           `mapstructure:"domain_name"`
	Port          int              `mapstructure:"port"`
	BaseUriPrefix string           `mapstructure:"base_uri_prefix"`
	AuthTokens    AuthTokensConfig `mapstructure:"auth_tokens"`
	Database      DatabaseConfig   `mapstructure:"database"`
}

func loadConfig(configPath string) (*Config, error) {
	viper.SetConfigName(strings.TrimSuffix(filepath.Base(configPath), filepath.Ext(configPath)))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(configPath))
	configureEnvVarsAndDefaults()

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

func validateConfig() error {
	errorMessageTemplate := "no '%s' keyword found in config"

	if err := validateNestedYaml(errorMessageTemplate, "database", []string{"host", "user", "password", "dbname", "sslmode"}); err != nil {
		return err
	}

	if err := validateNestedYaml(errorMessageTemplate, "auth_tokens", []string{"jwt_secret_key"}); err != nil {
		return err
	}

	if err := validateSingleKeyword(errorMessageTemplate, "domain_name"); err != nil {
		return err
	}

	return nil
}

func validateSingleKeyword(errorMessageTemplate string, keyword string) error {
	if !viper.IsSet(keyword) {
		return fmt.Errorf(errorMessageTemplate, keyword)
	}
	return nil
}

func validateNestedYaml(errorMessageTemplate string, topLevelKeyword string, requiredChildren []string) error {
	if !viper.IsSet(topLevelKeyword) {
		return fmt.Errorf(errorMessageTemplate, topLevelKeyword)
	}
	for _, child := range requiredChildren {
		keywordFullName := fmt.Sprintf("%s.%s", topLevelKeyword, child)
		if !viper.IsSet(keywordFullName) {
			return fmt.Errorf(errorMessageTemplate, keywordFullName)
		}
	}
	return nil
}
