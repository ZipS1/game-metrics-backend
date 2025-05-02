package repository

import (
	"game-metrics/game-service/internal/models"
)

func CreatePlayer(playerId, activityId uint) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}

	player := models.Player{Id: playerId, ActivityId: activityId}
	if result := db.Create(&player); result.Error != nil {
		return result.Error
	}

	return nil
}
