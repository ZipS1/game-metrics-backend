package amqp_handlers

import (
	"encoding/json"
	"game-metrics/game-service/internal/repository"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

func handleActivityCreated(delivery amqp091.Delivery, logger zerolog.Logger) {
	var message struct {
		ID     uint   `json:"id"`
		UserId string `json:"user-id"`
	}

	if err := json.Unmarshal(delivery.Body, &message); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal message")
		return
	}

	userUuid, err := uuid.Parse(message.UserId)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse UUID")
		return
	}

	if err := repository.CreateActivity(message.ID, userUuid); err != nil {
		logger.Error().Err(err).Msg("Failed to store activity in database")
		return
	}

	logger.Info().Uint("activity-id", message.ID).Msg("Activity saved")
}
