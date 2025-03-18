package middlewares

import (
	"crypto/ed25519"
	"errors"
	"game-metrics/auth-service/internal/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type PublicKeyProvider interface {
	GetPublicKey() (ed25519.PublicKey, error)
}

func RequireAuth(provider PublicKeyProvider) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getJwt(ctx)
		if err != nil {
			abortUnauthorized(ctx, err)
			return
		}

		key, err := provider.GetPublicKey()
		if err != nil {
			abortInternalError(ctx, err)
			return
		}

		if err := jwt.ValidateToken(token, key); err != nil {
			abortUnauthorized(ctx, err)
			return
		}

		ctx.Next()
	}
}

func abortUnauthorized(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	ctx.Abort()
}

func abortInternalError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	ctx.Abort()
}

func getJwt(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		cookieValue, err := ctx.Cookie("access_token")
		if err != nil {
			return "", err
		}

		return cookieValue, nil

	}

	parts := strings.Fields(authHeader)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("authorization header format must be Bearer {token}")
	}

	return parts[1], nil
}
