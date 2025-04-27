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

	activityId, err := repository.CreateDefaultActivityForUser(uuid)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create default activity in database")
		return
	}

	amqp.SendMessage("activity created", map[string]interface{}{
		"id":      activityId,
		"user-id": message.UserId,
	}, logger)
	logger.Info().Str("user-id", message.UserId).Msg("Default activity created for user")
}
