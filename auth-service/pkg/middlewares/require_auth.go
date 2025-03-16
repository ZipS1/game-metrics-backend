package middlewares

import "github.com/gin-gonic/gin"

func RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
