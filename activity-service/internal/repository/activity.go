package repository

import (
	"fmt"
	"game-metrics/activity-service/internal/models"

	"github.com/google/uuid"
)

func CreateDefaultActivityForUser(userId uuid.UUID) (uint, error) {
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

func GetUserActivities(userId uuid.UUID) ([]models.Activity, error) {
	db, err := connectToDatabase()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	var activities []models.Activity
	result := db.Where("user_id = ?", userId).Find(&activities)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get activities from database: %w", err)
	}

	return activities, nil
}
