package controllers

import (
	"errors"
	"game-metrics/activity-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func GetActivities(logger zerolog.Logger) gin.HandlerFunc {
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

		activities, err := repository.GetUserActivities(userId)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to get user activities", logger)
			return
		}

		var response []map[string]interface{}
		for _, activity := range activities {
			response = append(response, map[string]interface{}{
				"id":   activity.ID,
				"name": activity.Name,
			})
		}

		ctx.JSON(http.StatusOK, response)
		logger.Info().Str("user-id", userId.String()).Msg("Activities sent successfully")
	}
}

func failWithError(ctx *gin.Context, err error, code int, msg string, logger zerolog.Logger) {
	logger.Error().Err(err).Msg(msg)
	ctx.AbortWithStatusJSON(code, gin.H{
		"error": err.Error(),
	})
}
