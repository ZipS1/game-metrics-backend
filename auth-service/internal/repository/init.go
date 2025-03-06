package repository

import (
	"game-metrics/auth-service/internal/models"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
	initErr    error
)

func Init(dsn string) error {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			initErr = err
			return
		}
		if err := db.AutoMigrate(&models.User{}); err != nil {
			initErr = err
			return
		}

		dbInstance = db
	})

	return initErr
}
