package dto

type DeltaGamePlayerDTO struct {
	Id          uint `json:"id" binding:"required"`
	PointsDelta int  `json:"delta" binding:"required"`
}
