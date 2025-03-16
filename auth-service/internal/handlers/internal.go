package handlers

import (
	"game-metrics/auth-service/internal/config"
	"game-metrics/auth-service/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func configureInternalEndpoints(r *gin.RouterGroup, config config.Config, logger zerolog.Logger) {
	r.GET("/jwks", controllers.Jwks(config.JwtToken.Ed25519PublicKey, logger))
}
