package dto

type CreateGamePlayerDTO struct {
	Id          uint `json:"id" binding:"required"`
	EntryPoints uint `json:"entryPoints" binding:"required"`
}
