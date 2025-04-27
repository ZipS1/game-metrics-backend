package repository

import (
	"game-metrics/players-service/internal/models"

	"github.com/google/uuid"
)

func CreateActivity(id uint, userId uuid.UUID) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}

	activity := models.Activity{ID: id, UserID: userId}
	if result := db.Create(&activity); result.Error != nil {
		return result.Error
	}

	return nil
}
