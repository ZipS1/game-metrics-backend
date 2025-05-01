package amqp_handlers

import (
	"game-metrics/game-service/internal/config"

	"github.com/rs/zerolog"
)

func ConfigureHandlers(amqpConfig config.AMQPConfig, logger zerolog.Logger) {
	logger.Info().Msg("AMQP handlers started")
}
