package handlers

import (
	"game-metrics/auth-service/internal/config"
	"game-metrics/auth-service/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func configureApiEndpoints(r *gin.RouterGroup, authTokensConfig config.AuthTokensConfig, logger zerolog.Logger) {
	r.POST("/register", controllers.Register(logger))
	r.POST("/login", controllers.Login(authTokensConfig.JwtExpirationTime, logger))
}
