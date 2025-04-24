package repository

import "game-metrics/players-service/internal/models"

func CreateActivity(ID uint) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}

	activity := models.Activity{ID: ID}
	if result := db.Create(&activity); result.Error != nil {
		return result.Error
	}

	return nil
}
