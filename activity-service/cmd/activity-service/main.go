package main

import (
	"fmt"
	"game-metrics/activity-service/internal/amqp"
	"game-metrics/activity-service/internal/config"
	"game-metrics/activity-service/internal/handlers"
	"game-metrics/activity-service/internal/repository"
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

	if err := repository.Init(cfg.Database.GetConnectionString()); err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	closeConn, err := amqp.Init(cfg.AMQP.GetConnectionString(), cfg.AMQP.Timeout)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to message broker")
	}
	defer closeConn()

	handlers.ConfigureRouter(r, *cfg, logger)

	var port string = fmt.Sprintf(":%d", cfg.Port)
	if err := r.Run(port); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start Activity service")
	}
}
