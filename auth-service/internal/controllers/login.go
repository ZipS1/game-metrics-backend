package controllers

import (
	"fmt"
	"game-metrics/auth-service/internal/config"
	"game-metrics/auth-service/internal/jwt"
	"game-metrics/auth-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

func Login(config config.Config, logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8"`
		}
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			respondWithError(ctx, err, http.StatusBadRequest, "Incorrect JSON passed", logger)
			return
		}

		user, err := repository.GetUserByEmail(requestBody.Email)
		if err != nil {
			respondWithError(ctx, err, http.StatusUnauthorized, "Invalid email or password", logger)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(requestBody.Password)); err != nil {
			respondWithError(ctx, err, http.StatusUnauthorized, "Invalid email or password", logger)
			return
		}

		jwtToken, err := jwt.GenerateNewTokenForUser(
			*user,
			config.JwtToken.JwtExpirationTime,
			config.JwtToken.Ed25519PrivateKey,
		)
		if err != nil {
			respondWithError(ctx, err, http.StatusInternalServerError, "Failed to generate access token", logger)
			return
		}

		respondWithAccessToken(ctx, http.StatusOK, jwtToken)
		logger.Info().Msg(fmt.Sprintf("User %s successfully logged in with token: %s", user.Email, jwtToken))
	}
}
