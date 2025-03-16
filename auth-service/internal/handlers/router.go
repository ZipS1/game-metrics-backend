package handlers

import (
	"game-metrics/auth-service/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ConfigureRouter(r *gin.Engine, config config.Config, logger zerolog.Logger) {
	internalRouter := r.Group("/internal")
	configureInternalEndpoints(internalRouter, config.JwtToken, logger)

	publicRouter := r.Group(config.BaseUriPrefix)
	configureHealthEndpoint(publicRouter, logger)
	configureApiEndpoints(publicRouter, config, logger)
}
