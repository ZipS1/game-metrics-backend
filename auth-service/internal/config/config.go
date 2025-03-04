package config

import (
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Config struct {
	Port          int    `mapstructure:"port"            json:"port"`
	BaseUriPrefix string `mapstructure:"base_uri_prefix" json:"base_uri_prefix"`
}

func LoadConfig(configPath string, logger zerolog.Logger) (*Config, error) {
	defaults := Config{
		Port:          8080,
		BaseUriPrefix: "/api/auth",
	}

	if configPath == "" {
		logger.Info().Interface("defaults", defaults).Msg("Using default values")
	} else {
		viper.SetConfigName(strings.TrimSuffix(filepath.Base(configPath), filepath.Ext(configPath)))
		viper.SetConfigType("yaml")
		viper.AddConfigPath(filepath.Dir(configPath))
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	viper.AutomaticEnv()
	setDefaults(defaults)

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setDefaults(defaults Config) {
	viper.SetDefault("port", defaults.Port)
	viper.SetDefault("path_prefix", defaults.BaseUriPrefix)
}
