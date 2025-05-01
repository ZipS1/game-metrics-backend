package api_handlers

import (
	"game-metrics/libs/auth_middleware"
	"game-metrics/players-service/internal/config"
	"game-metrics/players-service/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func configureApiEndpoints(r *gin.RouterGroup, config config.Config, logger zerolog.Logger) {
	publicKeyProvider := PublicKeyProvider{}
	publicKeyProvider.Init(config.JwksEndpoint)

	r.POST("/", auth_middleware.RequireAuth(publicKeyProvider, logger), controllers.CreatePlayer(logger))
}
