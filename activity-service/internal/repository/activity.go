package repository

import (
	"fmt"
	"game-metrics/activity-service/internal/models"

	"github.com/google/uuid"
)

func CreateActivity(userId uuid.UUID, name string) (uint, error) {
	db, err := connectToDatabase()
	if err != nil {
		return 0, err
	}

	activity := models.Activity{UserId: userId, Name: name}
	if result := db.Create(&activity); result.Error != nil {
		return 0, result.Error
	}

	return activity.ID, nil
}

func CreateDefaultActivity(userId uuid.UUID) (uint, error) {
	return CreateActivity(userId, "Активность 1")
}

func GetUserActivities(userId uuid.UUID) ([]models.Activity, error) {
	db, err := connectToDatabase()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	var activities []models.Activity
	result := db.Where("user_id = ?", userId).Find(&activities)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get activities from database: %w", result.Error)
	}

	return activities, nil
}
