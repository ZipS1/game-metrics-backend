package controllers

import (
	"errors"
	"game-metrics/players-service/internal/amqp"
	"game-metrics/players-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func CreatePlayer(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody struct {
			ActivityId uint   `json:"activityId" binding:"required"`
			Name       string `json:"name" binding:"required"`
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

		if err := repository.ValidateActivityAccess(userId, requestBody.ActivityId); err != nil {
			failWithError(ctx, err, http.StatusForbidden, "Activity does not exist or you have no access to it", logger)
			return
		}

		playerId, err := repository.CreatePlayer(userId, requestBody.ActivityId, requestBody.Name)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to create player", logger)
			return
		}

		if err = amqp.SendMessage("player-created", map[string]any{
			"activityId": requestBody.ActivityId,
			"playerId":   playerId,
		}, logger); err != nil {
			logger.Error().Err(err).Uint("player-id", playerId).Msg("Failed to send player-created amqp message")
		}

		respondWithSuccess(ctx, http.StatusCreated, "Player successfully created", logger)
	}
}
