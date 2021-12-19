package model

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	gorm.Model
	ID        uint
	PeepID    uint
	Peep      Peep
	CreatedAt time.Time
	UpdatedAt time.Time
	Note      string
}
