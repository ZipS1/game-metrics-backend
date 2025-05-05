package main

import (
	"fmt"
	"game-metrics/auth-service/internal/amqp"
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

	cfg, err := config.ConstructConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to load configuration")
	}

	closeRepositoryFunc, err := repository.Init(cfg.Database.GetConnectionString())
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to the database")
	}
	defer closeRepositoryFunc()

	closeAmqpFunc, err := amqp.Init(cfg.AMQP.GetConnectionString(), cfg.AMQP.Timeout)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to message broker")
	}
	defer closeAmqpFunc()

	handlers.ConfigureRouter(r, *cfg, logger)

	var port string = fmt.Sprintf(":%d", cfg.Port)
	if err := r.Run(port); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start Auth Service")
	}
}
