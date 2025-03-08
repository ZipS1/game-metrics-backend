package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Login(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"login": "hello",
		})
	}
}
