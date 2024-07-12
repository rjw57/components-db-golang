package models

import (
	"time"

	"gorm.io/gorm"
)

type Cabinet struct {
	gorm.Model
	Id        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
