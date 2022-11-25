package entities

import (
	"gorm.io/gorm"
	"time"
)

type Entity struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
