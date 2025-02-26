package config

import "github.com/spf13/viper"

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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
