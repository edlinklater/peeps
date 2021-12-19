package model

import (
	"gorm.io/gorm"
	"time"
)

type Peep struct {
	gorm.Model
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	RealName  string
}
