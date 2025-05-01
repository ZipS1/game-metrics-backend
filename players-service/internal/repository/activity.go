package repository

import (
	"errors"
	"fmt"
	"game-metrics/players-service/internal/models"

	"github.com/google/uuid"
)

var ErrActivityAccessDenied = errors.New("user does not have access to this activity")

func CreateActivity(id uint, userId uuid.UUID) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}

	activity := models.Activity{Id: id, UserId: userId}
	if result := db.Create(&activity); result.Error != nil {
		return result.Error
	}

	return nil
}

func ValidateActivityAccess(userId uuid.UUID, activityId uint) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}

	var activity models.Activity
	result := db.Where("id = ?", activityId).First(&activity)
	if result.Error != nil {
		return fmt.Errorf("failed to get activity from database: %w", result.Error)
	}

	if userId != activity.UserId {
		return ErrActivityAccessDenied
	}

	return nil
}
