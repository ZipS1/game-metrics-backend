package controllers

import (
	"fmt"
	"game-metrics/auth-service/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Jwks(jwtConfig config.JwtTokenConfig, logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"alg":  "Ed25519",
			"jwks": jwtConfig.Ed25519PublicKey,
		})
		logger.Info().Msg(fmt.Sprintf("Jwt endpoint reached by %s", ctx.ClientIP()))
	}
}
