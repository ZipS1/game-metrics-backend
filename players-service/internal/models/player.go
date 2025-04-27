package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	ActivityId uint   `gorm:"not null;uniqueIndex:idx_activity_name"`
	Name       string `gorm:"not null;uniqueIndex:idx_activity_name"`
	Score      int    `gorm:"not null"`

	Activity Activity `gorm:"foreignKey:ActivityId;references:Id"`
}
