package handlers

import (
	"game-metrics/auth-service/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ConfigureRouter(r *gin.Engine, config config.Config, logger zerolog.Logger) {
	baseRouter := r.Group(config.BaseUriPrefix)
	configureHealthEndpoint(baseRouter, logger)
	configureApiEndpoints(baseRouter, config, logger)
}
