package main

import (
	"game-metrics/api-gateway/internal/handlers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	r := gin.Default()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	handlers.ConfigureHealthEndpoint(r, logger)
	handlers.ConfigureApiEndpoints(r, logger)

	if err := r.Run(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start API Gateway")
	}
}
