package repository

import (
	"game-metrics/auth-service/internal/models"
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

func GetUserHashedPasswordByEmail(email string) (string, error) {
	db, err := connectToDatabase()
	if err != nil {
		return "", err
	}

	var user models.User
	if result := db.First(&user, "email = ?", email); result.Error != nil {
		return "", result.Error
	}
	return user.PasswordHash, nil
}
