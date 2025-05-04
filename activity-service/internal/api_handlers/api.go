package api_handlers

import (
	"game-metrics/activity-service/internal/config"
	"game-metrics/activity-service/internal/controllers"

	"game-metrics/libs/auth_middleware"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func configureApiEndpoints(r *gin.RouterGroup, config config.Config, logger zerolog.Logger) {
	publicKeyProvider := PublicKeyProvider{}
	publicKeyProvider.Init(config.JwksEndpoint)

	r.GET("/", auth_middleware.RequireAuth(publicKeyProvider, logger), controllers.GetActivities(logger))
	r.POST("/", auth_middleware.RequireAuth(publicKeyProvider, logger), controllers.CreateActivity(logger))
}
