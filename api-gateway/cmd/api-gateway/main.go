package main

import (
	"fmt"
	"game-metrics/api-gateway/internal/config"
	"game-metrics/api-gateway/internal/handlers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	r := gin.Default()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load configuration")
	}

	handlers.ConfigureHealthEndpoint(r, logger)
	handlers.ConfigureApiEndpoints(r, logger, cfg.Services)

	var port string = fmt.Sprintf(":%d", cfg.Port)
	if err := r.Run(port); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start API Gateway")
	}
}
