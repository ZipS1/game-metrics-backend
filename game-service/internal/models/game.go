package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Duration   time.Duration
	ActivityId uint     `gorm:"not null"`
	Players    []Player `gorm:"many2many:game_players;"`

	Activity Activity `gorm:"foreignKey:ActivityId;references:Id"`
}
