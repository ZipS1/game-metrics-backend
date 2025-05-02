package repository

import (
	"game-metrics/game-service/internal/dto"
	"game-metrics/game-service/internal/models"
)

func CreateGame(activityId uint, players []dto.CreateGamePlayerDTO) (uint, error) {
	db, err := connectToDatabase()
	if err != nil {
		return 0, err
	}

	game := models.Game{ActivityId: activityId}
	if result := db.Create(&game); result.Error != nil {
		return 0, result.Error
	}

	for _, player := range players {
		player := models.GamePlayer{
			GameID:      game.ID,
			PlayerID:    player.Id,
			EntryPoints: player.EntryPoints,
		}

		if result := db.Create(&player); result.Error != nil {
			return 0, result.Error
		}
	}

	return game.ID, nil
}
