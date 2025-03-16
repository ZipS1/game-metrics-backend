package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func CheckAuth(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		respondWithSuccess(ctx, http.StatusOK, "You are logged in", logger)
	}
}
