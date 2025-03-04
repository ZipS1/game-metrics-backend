package main

import (
	"fmt"
	"game-metrics/auth-service/internal/config"
	"game-metrics/auth-service/internal/handlers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	r := gin.Default()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	args := config.ParseArgs(logger)
	cfg, err := config.LoadConfig(args.ConfigPath, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load configuration")
	}

	baseRouter := r.Group(cfg.BaseUriPrefix)
	handlers.ConfigureHealthEndpoint(baseRouter, logger)
	handlers.ConfigureApiEndpoints(baseRouter, logger)

	var port string = fmt.Sprintf(":%d", cfg.Port)
	if err := r.Run(port); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start Auth Service")
	}
}
