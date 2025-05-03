package amqp_handlers

import (
	"encoding/json"
	"game-metrics/players-service/internal/dto"
	"game-metrics/players-service/internal/repository"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

func handleGameFinished(delivery amqp091.Delivery, logger zerolog.Logger) {
	var message struct {
		GameId uint                     `json:"gameId"`
		Deltas []dto.DeltaGamePlayerDTO `json:"players" binding:"required,dive"`
	}

	if err := json.Unmarshal(delivery.Body, &message); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal message")
		return
	}

	if err := repository.UpdatePlayerScores(message.Deltas); err != nil {
		logger.Error().Err(err).Msg("Failed to store update player scores")
		return
	}

	logger.Info().Uint("game-id", message.GameId).Msg("Game players score updated")
}
