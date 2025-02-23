package main

import (
	"game-metrics/api-gateway/pkg/rproxy"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	r := gin.Default()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	r.GET("/health", func(c *gin.Context) {
		logger.Info().
			Str("Client IP", c.ClientIP()).
			Str("Proxy-Service", "api-gateway").
			Str("Endpoint", c.Request.URL.RawPath).
			Send()

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	api := r.Group("/api")
	{
		api.Any("/auth/*proxyPath", rproxy.ReverseProxy("http://auth-service:8080"), func(c *gin.Context) {
			logger.
				Info().
				Str("Client IP", c.ClientIP()).
				Str("Proxy-Service", "auth").
				Str("Endpoint", c.Request.URL.RequestURI()).
				Send()
		})
	}

	r.Run()
}
