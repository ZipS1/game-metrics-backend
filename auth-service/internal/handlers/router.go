package handlers

import (
	"game-metrics/auth-service/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ConfigureRouter(r *gin.Engine, baseUriPrefix string, authTokensConfig config.AuthTokensConfig, logger zerolog.Logger) {
	baseRouter := r.Group(baseUriPrefix)
	configureHealthEndpoint(baseRouter, logger)
	configureApiEndpoints(baseRouter, authTokensConfig, logger)
}
