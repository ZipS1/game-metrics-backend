package amqp_handlers

import (
	"game-metrics/activity-service/internal/amqp"
	"game-metrics/activity-service/internal/config"

	"github.com/rs/zerolog"
)

func ConfigureHandlers(amqpConfig config.AMQPConfig, logger zerolog.Logger) {
	if err := amqp.RunHandler("auth-service", "activity-service", "user-created", amqpConfig.Timeout, handleUserCreated, logger); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run user-created handler")
	}

	logger.Info().Msg("AMQP handlers started")
}
