package repository

import (
	"game-metrics/auth-service/internal/models"

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

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	connectionString = connStr
	return nil
}
