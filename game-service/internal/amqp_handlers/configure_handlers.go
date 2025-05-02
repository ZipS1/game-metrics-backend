package amqp_handlers

import (
	"game-metrics/game-service/internal/amqp"
	"game-metrics/game-service/internal/config"

	"github.com/rs/zerolog"
)

func ConfigureHandlers(amqpConfig config.AMQPConfig, logger zerolog.Logger) {
	if err := amqp.RunHandler("activity-service", "game-service-activity-created", "activity created", amqpConfig.Timeout, handleActivityCreated, logger); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run activity created handler")
	}

	if err := amqp.RunHandler("players-service", "game-service-player-created", "player-created", amqpConfig.Timeout, handlePlayerCreated, logger); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run player-created handler")
	}

	logger.Info().Msg("AMQP handlers started")
}
