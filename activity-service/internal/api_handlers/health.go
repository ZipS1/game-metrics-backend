package api_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func configureHealthEndpoint(r *gin.RouterGroup, logger zerolog.Logger) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
		logger.Info().Msg("Healthcheck endpoint reached")
	})
}
