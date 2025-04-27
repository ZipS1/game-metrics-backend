package models

import "github.com/google/uuid"

type Activity struct {
	Id     uint      `gorm:"primaryKey"`
	UserId uuid.UUID `gorm:"type:uuid;not null"`
}
