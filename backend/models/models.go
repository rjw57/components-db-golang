package models

import (
	"time"

	"github.com/google/uuid"
)

type Cabinet struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `gorm:"index;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
