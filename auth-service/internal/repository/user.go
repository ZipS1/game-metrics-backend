package repository

import (
	"errors"
	"game-metrics/auth-service/internal/models"

	"gorm.io/gorm"
)

func CreateUser(email string, passwordHash string) (string, error) {
	db, err := connectToDatabase()
	if err != nil {
		return "", err
	}

	user := models.User{Email: email, PasswordHash: passwordHash}
	if result := db.Create(&user); result.Error != nil {
		return "", result.Error
	}

	return user.ID.String(), nil
}

func GetUserByEmail(email string) (*models.User, error) {
	db, err := connectToDatabase()
	if err != nil {
		return nil, err
	}

	var user models.User
	if result := db.First(&user, "email = ?", email); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UserIdExists(userId string) (bool, error) {
	db, err := connectToDatabase()
	if err != nil {
		return false, err
	}

	var user models.User
	result := db.First(&user, "id = ?", userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
