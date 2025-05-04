package controllers

import (
	"errors"
	"game-metrics/activity-service/internal/amqp"
	"game-metrics/activity-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func CreateActivity(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody struct {
			Name string `json:"name" binding:"required"`
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

		activityId, err := repository.CreateActivity(userId, requestBody.Name)
		if err != nil {
			failWithError(ctx, err, http.StatusInternalServerError, "Failed to create activity", logger)
			return
		}

		if err = amqp.SendMessage("activity created", map[string]any{
			"activityId": activityId,
			"user-id":    userId,
		}, logger); err != nil {
			logger.Error().Err(err).Uint("activity-id", activityId).Msg("Failed to send activity created amqp message")
		}

		respondWithSuccess(ctx, http.StatusCreated, "Activity successfully created", logger)
	}
}
