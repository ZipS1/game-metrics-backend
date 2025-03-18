package controllers

import (
	"crypto/ed25519"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Jwks(key ed25519.PublicKey, logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"alg":  "Ed25519",
			"jwks": key,
		})
		logger.Info().Msg(fmt.Sprintf("Jwt endpoint reached by %s", ctx.ClientIP()))
	}
}
