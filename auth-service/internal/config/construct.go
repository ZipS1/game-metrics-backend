package config

import "github.com/rs/zerolog"

func ConstructConfig(logger zerolog.Logger) (*Config, error) {
	args := parseArgs(logger)
	cfg, err := loadConfig(args.ConfigPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
