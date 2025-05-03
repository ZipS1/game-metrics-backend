package amqp_handlers

import (
	"encoding/json"
	"game-metrics/game-service/internal/repository"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

func handlePlayerCreated(delivery amqp091.Delivery, logger zerolog.Logger) {
	var message struct {
		ActivityId uint `json:"activityId"`
		PlayerId   uint `json:"playerId"`
	}

	if err := json.Unmarshal(delivery.Body, &message); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal message")
		return
	}

	if err := repository.CreatePlayer(message.PlayerId, message.ActivityId); err != nil {
		logger.Error().Err(err).Msg("Failed to store player in database")
		return
	}

	logger.Info().Uint("player-id", message.PlayerId).Uint("activity-id", message.ActivityId).Msg("Player saved")
}
