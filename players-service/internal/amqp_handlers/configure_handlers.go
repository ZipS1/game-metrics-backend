package amqp_handlers

import (
	"game-metrics/players-service/internal/amqp"
	"game-metrics/players-service/internal/config"

	"github.com/rs/zerolog"
)

func ConfigureHandlers(amqpConfig config.AMQPConfig, logger zerolog.Logger) {
	if err := amqp.RunHandler("activity-service", "players-service-activity-created", "activity created", amqpConfig.Timeout, handleActivityCreated, logger); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run activity-created handler")
	}

	if err := amqp.RunHandler("game-service", "players-service-game-finished", "game-finished", amqpConfig.Timeout, handleGameFinished, logger); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run game-finished handler")
	}

	logger.Info().Msg("AMQP handlers started")
}
