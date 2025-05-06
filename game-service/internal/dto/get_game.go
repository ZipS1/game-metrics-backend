package dto

import "time"

type GetGameDTO struct {
	ID        uint               `json:"id"`
	StartTime time.Time          `json:"startTime"`
	Duration  time.Duration      `json:"duration"`
	Players   []GetGamePlayerDTO `json:"players"`
}
