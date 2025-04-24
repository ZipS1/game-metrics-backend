package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	ActivityId uint
	Name       string
	Score      int
}
