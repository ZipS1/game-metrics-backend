package controllers

import (
	"fmt"
	"game-metrics/auth-service/internal/jwt"
	"game-metrics/auth-service/internal/repository"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

func Login(jwtExpirationTime time.Duration, logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8"`
		}
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			respondWithError(ctx, err, "Incorrect JSON passed", logger)
			return
		}

		user, err := repository.GetUserByEmail(requestBody.Email)
		if err != nil {
			respondWithError(ctx, err, "Invalid email or password", logger)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(requestBody.Password)); err != nil {
			respondWithError(ctx, err, "Invalid email or password", logger)
			return
		}

		jwtToken, err := jwt.GenerateNewTokenForUser(jwt.UserClaims{FirstName: user.FirstName, LastName: user.LastName}, jwtExpirationTime)
		if err != nil {
			respondWithError(ctx, err, "Failed to generate access token", logger)
		}

		ctx.SetCookie("access_token", jwtToken, int(jwtExpirationTime), "/", os.Getenv("DOMAIN_NAME"), true, true)
		respondWithSuccess(ctx, fmt.Sprintf("User %s successfully logged in", requestBody.Email), logger)
	}
}
