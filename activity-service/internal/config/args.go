package config

import (
	"errors"
	"flag"
)

type Args struct {
	ConfigPath string
}

func parseArgs() (*Args, error) {
	configPath := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	if *configPath == "" {
		return nil, errors.New("--config flag is required")
	}

	args := Args{*configPath}
	return &args, nil
}
