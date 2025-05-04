package amqp_handlers

import (
	"encoding/json"
	"game-metrics/activity-service/internal/amqp"
	"game-metrics/activity-service/internal/repository"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

func handleUserCreated(delivery amqp091.Delivery, logger zerolog.Logger) {
	var message struct {
		UserId string `json:"id"`
	}

	if err := json.Unmarshal(delivery.Body, &message); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal message")
		return
	}

	uuid, err := uuid.Parse(message.UserId)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse UUID")
		return
	}

	activityId, err := repository.CreateDefaultActivity(uuid)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create default activity in database")
		return
	}

	if err = amqp.SendMessage("activity created", map[string]any{
		"activityId": activityId,
		"user-id":    message.UserId,
	}, logger); err != nil {
		logger.Error().Err(err).Uint("activity-id", activityId).Msg("Failed to send activity created amqp message")
	}
	logger.Info().Str("user-id", message.UserId).Msg("Default activity created for user")
}
