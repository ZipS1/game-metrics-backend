package repository

import (
	"game-metrics/game-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	connectionString string
)

func Init(connStr string) error {
	db, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		return err
	}

	if err := db.SetupJoinTable(&models.Game{}, "Players", &models.GamePlayer{}); err != nil {
		return err
	}

	if err := db.SetupJoinTable(&models.Player{}, "Games", &models.GamePlayer{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.Game{}, &models.Activity{}, &models.Player{}, &models.GamePlayer{}); err != nil {
		return err
	}

	connectionString = connStr
	return nil
}
