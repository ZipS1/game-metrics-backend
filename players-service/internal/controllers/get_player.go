package controllers

import (
	"encoding/json"
	"errors"
	"game-metrics/players-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func GetPlayer(logger zerolog.Logger) gin.HandlerFunc {
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

		playerIDStr := ctx.Param("id")
		if playerIDStr == "" {
			failWithError(ctx, errors.New("failed to get player id"), http.StatusBadRequest,
				"failed to get player id", logger)
			return
		}

		id, err := strconv.ParseUint(playerIDStr, 10, 64)
		if err != nil {
			failWithError(ctx, errors.New("failed to parse player id"), http.StatusInternalServerError,
				"Failed to parse player id", logger)
			return
		}
		playerID := uint(id)

		if err := repository.ValidatePlayerAccess(userId, playerID); err != nil {
			failWithError(ctx, err, http.StatusForbidden, "Player does not exist or you have no access to it", logger)
			return
		}

		player, err := repository.GetPlayer(playerID)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to get player", logger)
			return
		}

		data, err := json.Marshal(player)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to parse response to JSON", logger)
			return
		}

		ctx.Data(http.StatusOK, "application/json", data)
		logger.Info().Str("user-id", userId.String()).Msg("Player sent successfully")
	}
}
