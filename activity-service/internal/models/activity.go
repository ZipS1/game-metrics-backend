package models

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID        int
	UserId    uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
