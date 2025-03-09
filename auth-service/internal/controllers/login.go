package controllers

import (
	"fmt"
	"game-metrics/auth-service/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

func Login(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8"`
		}
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			respondWithError(ctx, err, "Incorrect JSON passed", logger)
			return
		}

		hash, err := repository.GetUserHashedPasswordByEmail(requestBody.Email)
		if err != nil {
			respondWithError(ctx, err, "Invalid email or password", logger)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(requestBody.Password)); err != nil {
			respondWithError(ctx, err, "Invalid email or password", logger)
			return
		}

		respondWithSuccess(ctx, fmt.Sprintf("User %s successfully logged in", requestBody.Email), logger)
	}
}
