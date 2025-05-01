package api_handlers

import (
	"game-metrics/game-service/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func configureApiEndpoints(r *gin.RouterGroup, config config.Config, logger zerolog.Logger) {
	publicKeyProvider := PublicKeyProvider{}
	publicKeyProvider.Init(config.JwksEndpoint)
}
