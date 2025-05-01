package repository

import (
	"game-metrics/players-service/internal/models"

	"github.com/google/uuid"
)

func CreatePlayer(userId uuid.UUID, activityId uint, name string) (uint, error) {
	db, err := connectToDatabase()
	if err != nil {
		return 0, err
	}

	player := models.Player{ActivityId: activityId, Name: name, Score: 0}
	if result := db.Create(&player); result.Error != nil {
		return 0, result.Error
	}

	return player.ID, nil
}
