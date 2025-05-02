package repository

import (
	"errors"
	"fmt"
	"game-metrics/game-service/internal/dto"
	"game-metrics/game-service/internal/models"

	"github.com/google/uuid"
)

var ErrGameAccessDenied = errors.New("user does not have access to this game")

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

func ValidateGameOwner(userId uuid.UUID, gameId uint) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}

	var game models.Game
	result := db.Preload("Activity").Where("id = ?", gameId).First(&game)
	if result.Error != nil {
		return fmt.Errorf("failed to get game from database: %w", result.Error)
	}

	if userId != game.Activity.UserId {
		return ErrGameAccessDenied
	}

	return nil
}

func AddPointsToPlayer(gameId, playerId, additionalPoints uint) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}

	var gamePlayer models.GamePlayer
	result := db.Where("game_id = ? AND player_id = ?", gameId, playerId).First(&gamePlayer)
	if result.Error != nil {
		return fmt.Errorf("failed to get game player from database: %w", result.Error)
	}

	gamePlayer.AdditionalPoints += additionalPoints

	saveResult := db.Save(&gamePlayer)
	if saveResult.Error != nil {
		return fmt.Errorf("failed to update additional points: %w", saveResult.Error)
	}

	return nil
}
