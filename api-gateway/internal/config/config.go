package config

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Name       string `mapstructure:"name"`
	PathPrefix string `mapstructure:"path_prefix"`
	Url        string `mapstructure:"url"`
}

type Config struct {
	Port     int             `mapstructure:"port"`
	Services []ServiceConfig `mapstructure:"services"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigName(strings.TrimSuffix(filepath.Base(configPath), filepath.Ext(configPath)))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(configPath))
	viper.AutomaticEnv()

	setDefaults()
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

func setDefaults() {
	viper.SetDefault("port", 8080)
}

func validateConfig() error {
	if !viper.IsSet("services") {
		return errors.New("no 'services' keyword in config found")
	}
	return nil
}
