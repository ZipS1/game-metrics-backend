package repository

import (
	"errors"
	"game-metrics/game-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init(connectionString string) (func(), error) {
	dbConn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}

	if err := dbConn.SetupJoinTable(&models.Game{}, "Players", &models.GamePlayer{}); err != nil {
		return nil, err
	}

	if err := dbConn.SetupJoinTable(&models.Player{}, "Games", &models.GamePlayer{}); err != nil {
		return nil, err
	}

	if err := dbConn.AutoMigrate(&models.Game{}, &models.Activity{}, &models.Player{}, &models.GamePlayer{}); err != nil {
		return nil, err
	}

	db = dbConn
	return close, err
}

func connectToDatabase() (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("repository is not initialized. Call Init() before using the database")
	}
	return db, nil
}

func close() {
	if db == nil {
		return
	}

	sqlDb, _ := db.DB()
	sqlDb.Close()
}
