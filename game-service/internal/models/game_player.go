package models

import "time"

type GamePlayer struct {
	GameID           uint `gorm:"primaryKey"`
	PlayerID         uint `gorm:"primaryKey"`
	EntryPoints      uint `gorm:"not null"`
	AdditionalPoints uint
	EndPoints        uint

	JoinedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
