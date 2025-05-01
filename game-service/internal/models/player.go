package models

type Player struct {
	Id         uint   `gorm:"primaryKey"`
	ActivityId uint   `gorm:"not null"`
	Games      []Game `gorm:"many2many:game_players;"`

	Activity Activity `gorm:"foreignKey:ActivityId;references:Id"`
}
