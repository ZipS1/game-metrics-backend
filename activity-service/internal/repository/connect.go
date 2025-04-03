package repository

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func isDatabaseInitialized() bool {
	return connectionString != ""
}

func connectToDatabase() (*gorm.DB, error) {
	if !isDatabaseInitialized() {
		return nil, errors.New("repository is not initialized. Call Init() before Connect()")
	}

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
