package repository

import (
	"game-metrics/players-service/internal/dto"
	"game-metrics/players-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func UpdatePlayerScores(deltas []dto.DeltaGamePlayerDTO) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for _, delta := range deltas {
			var player models.Player
			if result := tx.Where("id = ?", delta.Id).First(&player); result.Error != nil {
				return result.Error
			}

			player.Score -= delta.PointsDelta
			if result := tx.Save(&player); result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
}
