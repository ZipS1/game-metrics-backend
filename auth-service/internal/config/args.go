package config

import (
	"flag"
	"fmt"

	"github.com/rs/zerolog"
)

type Args struct {
	ConfigPath string
}

func parseArgs(logger zerolog.Logger) *Args {
	configPath := flag.String("config", "", "Path to the configuration file")
	flag.Parse()

	if *configPath == "" {
		logger.Fatal().Msg("Error: --config flag is required")
	}

	args := Args{*configPath}
	logger.Info().Msg(fmt.Sprintf("Using config file: %s\n", *configPath))
	return &args
}
