package dto

type GetPlayerDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}
