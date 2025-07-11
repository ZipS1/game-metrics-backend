package api_handlers

import (
	"game-metrics/game-service/internal/config"
	"game-metrics/game-service/internal/controllers"

	"game-metrics/libs/auth_middleware"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func configureApiEndpoints(r *gin.RouterGroup, config config.Config, logger zerolog.Logger) {
	publicKeyProvider := PublicKeyProvider{}
	publicKeyProvider.Init(config.JwksEndpoint)

	r.GET("/", auth_middleware.RequireAuth(publicKeyProvider, logger), controllers.GetGames(logger))
	r.GET("/:id", auth_middleware.RequireAuth(publicKeyProvider, logger), controllers.GetGame(logger))
	r.POST("/", auth_middleware.RequireAuth(publicKeyProvider, logger), controllers.CreateGame(logger))
	r.PATCH("/addPoints", auth_middleware.RequireAuth(publicKeyProvider, logger), controllers.AddPoints(logger))
	r.PUT("/finish", auth_middleware.RequireAuth(publicKeyProvider, logger), controllers.FinishGame(logger))
}
