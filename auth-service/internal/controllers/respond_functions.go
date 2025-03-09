package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func respondWithError(ctx *gin.Context, err error, message string, logger zerolog.Logger) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
	logger.Error().Err(err).Msg(message)
}

func respondWithSuccess(ctx *gin.Context, message string, logger zerolog.Logger) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": message,
	})
	logger.Info().Msg(message)
}
