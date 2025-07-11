package handlers

import (
	"crypto/ed25519"
	"game-metrics/auth-service/internal/config"
	"game-metrics/auth-service/internal/controllers"

	"game-metrics/libs/auth_middleware"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func configureApiEndpoints(r *gin.RouterGroup, config config.Config, logger zerolog.Logger) {
	publicKeyProviderMock := publicKeyProviderMock(func() (ed25519.PublicKey, error) {
		return config.JwtToken.Ed25519PublicKey, nil
	})

	r.POST("/register", controllers.Register(logger))
	r.POST("/login", controllers.Login(config, logger))
	r.GET("/check", auth_middleware.RequireAuth(publicKeyProviderMock, logger), controllers.CheckAuth(logger))
}
