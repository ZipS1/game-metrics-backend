package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ServiceProxyLogging(logger zerolog.Logger, serviceName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Info().
			Str("Client IP", ctx.ClientIP()).
			Str("Proxy-Service", serviceName).
			Str("Endpoint", ctx.Request.URL.RequestURI()).
			Send()
		ctx.Next()
	}
}
