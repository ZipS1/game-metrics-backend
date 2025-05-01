package amqp_handlers

import (
	"game-metrics/game-service/internal/amqp"
	"game-metrics/game-service/internal/config"

	"github.com/rs/zerolog"
)

func ConfigureHandlers(amqpConfig config.AMQPConfig, logger zerolog.Logger) {
	if err := amqp.RunHandler("activity-service", "game-service", "activity created", amqpConfig.Timeout, handleActivityCreated, logger); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run activity-created handler")
	}

	logger.Info().Msg("AMQP handlers started")
}
