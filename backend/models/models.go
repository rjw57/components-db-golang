package models

import (
	"time"

	"github.com/google/uuid"
)

type Cabinet struct {
	ID        uint      `fake:"skip"`
	UUID      uuid.UUID `fake:"skip" gorm:"index;type:uuid;default:gen_random_uuid();not null"`
	Name      string    `fake:"{adjective} {noun}" gorm:"not null"`
	CreatedAt time.Time `fake:"skip" gorm:"not null"`
	UpdatedAt time.Time `fake:"skip" gotm:"not null"`
}
