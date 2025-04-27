package models

import "github.com/google/uuid"

type Activity struct {
	ID     uint      `gorm:"primarykey"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
}
