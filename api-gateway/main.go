package main

import (
	"game-metrics/api-gateway/config"
	"game-metrics/api-gateway/pkg/rproxy"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func loggingMiddleware(logger zerolog.Logger, serviceName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Info().
			Str("Client IP", ctx.ClientIP()).
			Str("Proxy-Service", serviceName).
			Str("Endpoint", ctx.Request.URL.RequestURI()).
			Send()
		ctx.Next()
	}
}

func main() {
	r := gin.Default()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	r.GET("/health", loggingMiddleware(logger, "api-gateway"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	api := r.Group("/api")
	services := config.GetServices()
	for _, service := range services {
		serviceGroup := api.Group(service.PathPrefix)
		serviceGroup.Use(loggingMiddleware(logger, service.Name))
		serviceGroup.Any("/*servicePath", rproxy.ReverseProxy(service.URL))
	}

	if err := r.Run(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start API Gateway")
	}
}
