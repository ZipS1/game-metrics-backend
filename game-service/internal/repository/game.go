package repository

import (
	"errors"
	"fmt"
	"game-metrics/game-service/internal/dto"
	"game-metrics/game-service/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrGameAccessDenied = errors.New("user does not have access to this game")
var ErrGamePlayersNotFound = errors.New("not all game players found")
var ErrPointsMismatch = errors.New("points total does not match")

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

func GetGames(activityId uint) ([]dto.GetGamesGameDTO, error) {
	var gamesDTO []dto.GetGamesGameDTO

	db, err := connectToDatabase()
	if err != nil {
		return gamesDTO, err
	}

	var games []models.Game
	if result := db.Where("activity_id = ?", activityId).Find(&games); result.Error != nil {
		return gamesDTO, fmt.Errorf("failed to get games from database: %w", result.Error)
	}

	for _, game := range games {
		gameDTO := dto.GetGamesGameDTO{ID: game.ID, StartTime: game.CreatedAt, Duration: game.Duration}
		gamesDTO = append(gamesDTO, gameDTO)
	}

	return gamesDTO, nil
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

func ValidateFinishState(gameId uint, players []dto.FinishGamePlayerDTO) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}

	var gamePlayers []models.GamePlayer
	if err := db.Where("game_id = ?", gameId).Find(&gamePlayers).Error; err != nil {
		return fmt.Errorf("database error: %w", err)
	}

	gamePlayerIDs := make(map[uint]struct{})
	for _, gp := range gamePlayers {
		gamePlayerIDs[gp.PlayerID] = struct{}{}
	}

	dtoIDs := make(map[uint]struct{})
	for _, p := range players {
		dtoIDs[p.Id] = struct{}{}
	}

	if len(gamePlayerIDs) != len(dtoIDs) {
		return ErrGamePlayersNotFound
	}
	for id := range gamePlayerIDs {
		if _, exists := dtoIDs[id]; !exists {
			return ErrGamePlayersNotFound
		}
	}

	var totalEntryAdd, totalEnd uint
	for i := range gamePlayers {
		totalEntryAdd += gamePlayers[i].EntryPoints + gamePlayers[i].AdditionalPoints
	}
	for _, p := range players {
		totalEnd += p.EndPoints
	}

	if totalEntryAdd != totalEnd {
		return ErrPointsMismatch
	}

	return nil
}

func FinishGame(gameId uint, players []dto.FinishGamePlayerDTO) error {
	db, err := connectToDatabase()
	if err != nil {
		return err
	}

	var gamePlayers []models.GamePlayer
	if err := db.Where("game_id = ?", gameId).Find(&gamePlayers).Error; err != nil {
		return fmt.Errorf("failed to get game players: %w", err)
	}

	playerEndpoints := make(map[uint]uint)
	for _, p := range players {
		playerEndpoints[p.Id] = p.EndPoints
	}

	for i := range gamePlayers {
		if endpoints, exists := playerEndpoints[gamePlayers[i].PlayerID]; exists {
			gamePlayers[i].EndPoints = endpoints
		}
	}

	game, err := setGameDuration(db, gameId)
	if err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(game).Error; err != nil {
			return fmt.Errorf("failed to save game: %w", err)
		}
		for _, gp := range gamePlayers {
			if err := tx.Save(&gp).Error; err != nil {
				return fmt.Errorf("failed to save player %d: %w", gp.PlayerID, err)
			}
		}
		return nil
	})
}

func CalculatePlayerDelta(gameId uint) ([]dto.DeltaGamePlayerDTO, error) {
	var deltas []dto.DeltaGamePlayerDTO
	db, err := connectToDatabase()
	if err != nil {
		return deltas, err
	}

	var gamePlayers []models.GamePlayer
	if err := db.Where("game_id = ?", gameId).Find(&gamePlayers).Error; err != nil {
		return deltas, fmt.Errorf("failed to get game players: %w", err)
	}

	for _, player := range gamePlayers {
		delta := int(player.EndPoints) - int(player.EntryPoints+player.AdditionalPoints)
		deltas = append(deltas, dto.DeltaGamePlayerDTO{Id: player.PlayerID, PointsDelta: delta})
	}

	return deltas, nil
}

func setGameDuration(db *gorm.DB, gameId uint) (*models.Game, error) {
	var game models.Game
	if err := db.First(&game, gameId).Error; err != nil {
		return nil, fmt.Errorf("failed to get game: %w", err)
	}

	game.Duration = time.Since(game.CreatedAt)
	return &game, nil
}
