package repository

import (
	"game-metrics/activity-service/internal/models"

	"github.com/google/uuid"
)

func CreateDefaultActivityForUser(userId uuid.UUID) (int, error) {
	db, err := connectToDatabase()
	if err != nil {
		return 0, err
	}

	activity := models.Activity{UserId: userId, Name: "Активность 1"}
	if result := db.Create(&activity); result.Error != nil {
		return 0, result.Error
	}

	return activity.ID, nil
}
