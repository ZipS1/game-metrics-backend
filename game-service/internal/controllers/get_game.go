package controllers

import (
	"encoding/json"
	"errors"
	"game-metrics/game-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func GetGame(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIdValue, exists := ctx.Get("userId")
		if !exists {
			failWithError(ctx, errors.New("user ID not found in context"), http.StatusUnauthorized,
				"Missing authentication", logger)
			return
		}

		userId, err := uuid.Parse(userIdValue.(string))
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to get X-User-ID header", logger)
			return
		}

		gameIDStr := ctx.Param("id")
		if gameIDStr == "" {
			failWithError(ctx, errors.New("failed to get game id"), http.StatusBadRequest,
				"failed to get game id", logger)
			return
		}

		id, err := strconv.ParseUint(gameIDStr, 10, 64)
		if err != nil {
			failWithError(ctx, errors.New("failed to parse game id"), http.StatusInternalServerError,
				"Failed to parse game id", logger)
			return
		}
		gameID := uint(id)

		if err := repository.ValidateGameOwner(userId, gameID); err != nil {
			failWithError(ctx, err, http.StatusForbidden, "Game does not exist or you have no access to it", logger)
			return
		}

		game, err := repository.GetGame(gameID)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to get games", logger)
			return
		}

		data, err := json.Marshal(game)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to parse response to JSON", logger)
			return
		}

		ctx.Data(http.StatusOK, "application/json", data)
		logger.Info().Str("user-id", userId.String()).Msg("Games sent successfully")
	}
}
