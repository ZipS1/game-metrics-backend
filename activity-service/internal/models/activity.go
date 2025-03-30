package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	UserId uuid.UUID `gorm:"type:uuid;not null"`
	Name   string    `gorm:"not null"`
}
