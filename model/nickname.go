package model

import (
	"gorm.io/gorm"
	"time"
)

type Nickname struct {
	gorm.Model
	ID        uint
	PeepID    uint
	Peep      Peep
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Context   string
}
