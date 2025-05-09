package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func failWithError(ctx *gin.Context, err error, code int, msg string, logger zerolog.Logger) {
	logger.Error().Err(err).Msg(msg)
	ctx.AbortWithStatusJSON(code, gin.H{
		"error": err.Error(),
	})
}

func respondWithSuccess(ctx *gin.Context, code int, message string, logger zerolog.Logger) {
	ctx.JSON(code, gin.H{
		"message": message,
	})
	logger.Info().Msg(message)
}
