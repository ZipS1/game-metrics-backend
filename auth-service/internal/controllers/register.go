package controllers

import (
	"fmt"
	"game-metrics/auth-service/internal/amqp"
	"game-metrics/auth-service/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

		userId, err := repository.CreateUser(requestBody.Email, string(hash))
		if err != nil {
			switch err {
			case gorm.ErrDuplicatedKey:
				respondWithError(ctx, err, http.StatusConflict, "User with this email already registered", logger)
			default:
				respondWithError(ctx, err, http.StatusInternalServerError, "Failed to create user", logger)
			}
			return
		}

		amqp.SendMessage("user_registered", map[string]interface{}{
			"id": userId,
		}, logger)
		respondWithSuccess(ctx, http.StatusCreated, fmt.Sprintf("User %s successfully registered", requestBody.Email), logger)
	}
}
