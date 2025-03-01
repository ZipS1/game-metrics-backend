package handlers

import (
	"game-metrics/api-gateway/internal/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ConfigureHealthEndpoint(r *gin.Engine, logger zerolog.Logger) {
	r.GET("/health", middlewares.ServiceProxyLogging(logger, "api-gateway"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}
