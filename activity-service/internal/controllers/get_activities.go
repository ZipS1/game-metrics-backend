package controllers

import (
	"game-metrics/activity-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func GetActivities(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := uuid.Parse(ctx.GetHeader("X-User-ID"))
		if err != nil {
			logger.Error().Err(err).Msg("Failed to parse X-User-ID header")
			return
		}

		activities, err := repository.GetUserActivities(userId)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to get user activities")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve activities"})
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
