package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ConfigureApiEndpoints(r *gin.RouterGroup, logger zerolog.Logger) {
	r.GET("hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"login": "hello",
		})
	})
}
