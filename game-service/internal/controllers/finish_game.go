package controllers

import (
	"errors"
	"fmt"
	"game-metrics/game-service/internal/amqp"
	"game-metrics/game-service/internal/dto"
	"game-metrics/game-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func FinishGame(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody struct {
			GameId  uint                      `json:"gameId" binding:"required"`
			Players []dto.FinishGamePlayerDTO `json:"players" binding:"required,dive"`
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

		if err := repository.ValidateFinishState(requestBody.GameId, requestBody.Players); err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to validate finish state", logger)
			return
		}

		if err := repository.FinishGame(requestBody.GameId, requestBody.Players); err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to finish game", logger)
			return
		}

		deltas, err := repository.CalculatePlayerDelta(requestBody.GameId)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to finish game", logger)
			return
		}

		messagePayload := map[string]any{
			"gameId": requestBody.GameId,
		}

		for _, delta := range deltas {
			messagePayload[fmt.Sprintf("player-%d", delta.Id)] = delta.PointsDelta
		}

		if err = amqp.SendMessage("game-finished", messagePayload, logger); err != nil {
			logger.Error().Err(err).Uint("game-id", requestBody.GameId).Msg("Failed to send game-created amqp message")
		}

		ctx.JSON(http.StatusOK, messagePayload)
		logger.Info().Uint("game-id", requestBody.GameId).Msg("Game finished sucessfully")
	}
}
