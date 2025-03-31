package middlewares

import (
	"crypto/ed25519"
	"errors"
	"game-metrics/auth-service/internal/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type PublicKeyProvider interface {
	GetPublicKey() (ed25519.PublicKey, error)
}

func RequireAuth(provider PublicKeyProvider, logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getJwt(ctx)
		if err != nil {
			abortUnauthorized(ctx, err, logger)
			return
		}

		key, err := provider.GetPublicKey()
		if err != nil {
			abortInternalError(ctx, err, logger)
			return
		}

		userId, err := jwt.ValidateToken(token, key)
		if err != nil {
			abortUnauthorized(ctx, err, logger)
			return
		}

		ctx.Set("userId", userId)
		ctx.Next()
	}
}

func abortUnauthorized(ctx *gin.Context, err error, logger zerolog.Logger) {
	logger.Info().Err(err).Msg("Auth middleware failed: unauthorized")
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	ctx.Abort()
}

func abortInternalError(ctx *gin.Context, err error, logger zerolog.Logger) {
	logger.Info().Err(err).Msg("Auth middleware failed: internal error")
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	ctx.Abort()
}

func getJwt(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		cookieValue, err := ctx.Cookie("access_token")
		if err != nil {
			return "", errors.New("neither access_token cookie nor Authorization header is present")
		}

		return cookieValue, nil

	}

	parts := strings.Fields(authHeader)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("authorization header format must be Bearer {token}")
	}

	return parts[1], nil
}
