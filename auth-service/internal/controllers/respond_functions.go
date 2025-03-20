package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func respondWithError(ctx *gin.Context, err error, code int, message string, logger zerolog.Logger) {
	ctx.JSON(code, gin.H{
		"error": message,
	})
	logger.Error().Err(err).Msg(message)
}

func respondWithSuccess(ctx *gin.Context, code int, message string, logger zerolog.Logger) {
	ctx.JSON(code, gin.H{
		"message": message,
	})
	logger.Info().Msg(message)
}

func respondWithAccessToken(ctx *gin.Context, code int, access_token string) {
	ctx.JSON(code, gin.H{
		"message":      "Successfully logged in",
		"access_token": access_token,
	})
}
