package config

import (
	"flag"
	"fmt"

	"github.com/rs/zerolog"
)

type Args struct {
	ConfigPath string
}

func ParseArgs(logger zerolog.Logger) *Args {
	configPath := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	args := Args{*configPath}
	if *configPath == "" {
		logger.Info().Msg("No config file provided, default values will be set")
	} else {
		logger.Info().Msg(fmt.Sprintf("Using config file: %s\n", *configPath))
	}
	return &args
}
