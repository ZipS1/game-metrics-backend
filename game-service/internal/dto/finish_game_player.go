package dto

type FinishGamePlayerDTO struct {
	Id        uint `json:"id" binding:"required"`
	EndPoints uint `json:"endPoints" binding:"required"`
}
