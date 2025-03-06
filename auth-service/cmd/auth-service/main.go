package main

import (
	"fmt"
	"game-metrics/auth-service/internal/config"
	"game-metrics/auth-service/internal/handlers"
	"game-metrics/auth-service/internal/repository"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	r := gin.Default()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := config.ConstructConfig(logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load configuration")
	}

	if err := repository.Init(cfg.Database.GetConnectionString()); err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	handlers.ConfigureRouter(r, cfg.BaseUriPrefix, logger)

	var port string = fmt.Sprintf(":%d", cfg.Port)
	if err := r.Run(port); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start Auth Service")
	}
}
