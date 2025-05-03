package controllers

import (
	"errors"
	"game-metrics/game-service/internal/amqp"
	"game-metrics/game-service/internal/dto"
	"game-metrics/game-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func CreateGame(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody struct {
			ActivityId uint                      `json:"activityId" binding:"required"`
			Players    []dto.CreateGamePlayerDTO `json:"players" binding:"required,dive"`
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

		if err := repository.ValidateActivityOwner(userId, requestBody.ActivityId); err != nil {
			failWithError(ctx, err, http.StatusForbidden, "Activity does not exist or you have no access to it", logger)
			return
		}

		gameId, err := repository.CreateGame(requestBody.ActivityId, requestBody.Players)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to create game", logger)
			return
		}

		if err = amqp.SendMessage("game-created", map[string]any{
			"gameId": gameId,
		}, logger); err != nil {
			logger.Error().Err(err).Uint("game-id", gameId).Msg("Failed to send game-created amqp message")
		}

		respondWithSuccess(ctx, http.StatusCreated, "Game successfully created", logger)
	}
}
