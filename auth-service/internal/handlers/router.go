package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ConfigureRouter(r *gin.Engine, baseUriPrefix string, logger zerolog.Logger) {
	baseRouter := r.Group(baseUriPrefix)
	configureHealthEndpoint(baseRouter, logger)
	configureApiEndpoints(baseRouter, logger)
}
