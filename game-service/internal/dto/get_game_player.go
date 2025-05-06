package dto

type GetGamePlayerDTO struct {
	PlayerID         uint `json:"id"`
	EntryPoints      uint `json:"entryPoints"`
	AdditionalPoints uint `json:"additionalPoints"`
	EndPoints        uint `json:"endPoints"`
}
