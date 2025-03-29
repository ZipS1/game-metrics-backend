package handlers

import (
	"game-metrics/activity-service/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ConfigureRouter(r *gin.Engine, config config.Config, logger zerolog.Logger) {
	publicRouter := r.Group(config.PublicUriPrefix)
	configureHealthEndpoint(publicRouter, logger)
	configureApiEndpoints(publicRouter, config, logger)
}
