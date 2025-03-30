package api_handlers

import (
	"game-metrics/activity-service/internal/config"
	"game-metrics/activity-service/internal/controllers"
	"game-metrics/auth-service/pkg/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func configureApiEndpoints(r *gin.RouterGroup, config config.Config, logger zerolog.Logger) {
	publicKeyProvider := PublicKeyProvider{}
	publicKeyProvider.Init(config.JwksEndpoint)

	r.GET("/", middlewares.RequireAuth(publicKeyProvider, logger), controllers.GetActivities(logger))
}
