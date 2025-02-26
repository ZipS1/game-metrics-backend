package handlers

import (
	"game-metrics/api-gateway/config"
	"game-metrics/api-gateway/internal/middlewares"
	"game-metrics/api-gateway/pkg/rproxy"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ConfigureApiEndpoints(r *gin.Engine, logger zerolog.Logger) {
	api := r.Group("/api")
	services := config.GetServices()
	for _, service := range services {
		serviceGroup := api.Group(service.PathPrefix)
		serviceGroup.Use(middlewares.ServiceProxyLogging(logger, service.Name))
		serviceGroup.Any("/*servicePath", rproxy.ReverseProxy(service.URL))
	}
}
