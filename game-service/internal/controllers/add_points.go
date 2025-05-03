package controllers

import (
	"errors"
	"game-metrics/game-service/internal/amqp"
	"game-metrics/game-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func AddPoints(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody struct {
			GameId          uint `json:"gameId" binding:"required"`
			PlayerId        uint `json:"playerId" binding:"required"`
			AdditinalPoints uint `json:"additionalPoints" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			failWithError(ctx, err, http.StatusBadRequest, "Incorrect JSON passed", logger)
			return
		}

		userIdValue, exists := ctx.Get("userId")
		if !exists {
			failWithError(ctx, errors.New("user ID not found in context"), http.StatusUnauthorized,
				"Missing authentication", logger)
			return
		}

		userId, err := uuid.Parse(userIdValue.(string))
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to parse userId from context", logger)
			return
		}

		if err := repository.ValidateGameOwner(userId, requestBody.GameId); err != nil {
			failWithError(ctx, err, http.StatusForbidden, "Game does not exist or you have no access to it", logger)
			return
		}

		if err := repository.AddPointsToPlayer(requestBody.GameId, requestBody.PlayerId, requestBody.AdditinalPoints); err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to add points to game player", logger)
			return
		}

		if err = amqp.SendMessage("points-added", map[string]any{
			"gameId":       requestBody.GameId,
			"playerId":     requestBody.PlayerId,
			"points-added": requestBody.AdditinalPoints,
		}, logger); err != nil {
			logger.Error().Err(err).Uint("game-id", requestBody.GameId).Uint("player-id", requestBody.PlayerId).Int("points-added", int(requestBody.AdditinalPoints)).Msg("Failed to send game-created amqp message")
		}

		respondWithSuccess(ctx, http.StatusOK, "Points successfully added", logger)
	}
}
