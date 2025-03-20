package controllers

import (
	"fmt"
	"game-metrics/auth-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

func Register(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8"`
		}
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			respondWithError(ctx, err, http.StatusBadRequest, "Incorrect JSON passed", logger)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
		if err != nil {
			respondWithError(ctx, err, http.StatusInternalServerError, "Failed to hash password", logger)
			return
		}

		if _, err := repository.CreateUser(requestBody.Email, string(hash)); err != nil {
			switch err.Error() {
			case "duplicated key not allowed":
				respondWithError(ctx, err, http.StatusConflict, "User with this email already registered", logger)
			default:
				respondWithError(ctx, err, http.StatusInternalServerError, "Failed to create user", logger)
			}
			return
		}

		respondWithSuccess(ctx, http.StatusCreated, "Successfully registered", logger)
		logger.Info().Msg(fmt.Sprintf("User %s successfully registered", requestBody.Email))
	}
}
