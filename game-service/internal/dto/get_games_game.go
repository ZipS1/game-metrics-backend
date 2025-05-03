package dto

import "time"

type GetGamesGameDTO struct {
	ID        uint          `json:"id"`
	StartTime time.Time     `json:"startTime"`
	Duration  time.Duration `json:"duration"`
}
