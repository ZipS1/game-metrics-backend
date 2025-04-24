package amqp_handlers

import (
	"encoding/json"
	"game-metrics/players-service/internal/repository"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

func handleActivityCreated(delivery amqp091.Delivery, logger zerolog.Logger) {
	var message struct {
		ID uint `json:"id"`
	}

	if err := json.Unmarshal(delivery.Body, &message); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal message")
		return
	}

	err := repository.CreateActivity(message.ID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to store activity in database")
		return
	}

	logger.Info().Uint("activity-id", message.ID).Msg("Activity ID saved")
}
