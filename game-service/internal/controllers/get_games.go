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

func GetGames(logger zerolog.Logger) gin.HandlerFunc {
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

		activityIDStr := ctx.Query("activity_id")
		if activityIDStr == "" {
			failWithError(ctx, errors.New("activity_id query parameter is required"), http.StatusBadRequest,
				"Missing activity_id in query", logger)
			return
		}

		id, err := strconv.ParseUint(activityIDStr, 10, 64)
		if err != nil {
			failWithError(ctx, errors.New("activity_id query parameter is required"), http.StatusBadRequest,
				"Missing activity_id in query", logger)
			return
		}
		activityID := uint(id)

		if err := repository.ValidateActivityOwner(userId, activityID); err != nil {
			failWithError(ctx, err, http.StatusForbidden, "Activity does not exist or you have no access to it", logger)
			return
		}

		games, err := repository.GetGames(activityID)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to get games", logger)
			return
		}

		data, err := json.Marshal(games)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to parse response to JSON", logger)
			return
		}

		ctx.Data(http.StatusOK, "application/json", data)
		logger.Info().Str("user-id", userId.String()).Msg("Games sent successfully")
	}
}
