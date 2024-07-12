package models

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

type Cabinet struct {
	ID        uint      `fake:"skip"`
	UUID      uuid.UUID `fake:"skip" gorm:"index;type:uuid;default:gen_random_uuid();not null"`
	Name      string    `fake:"{adjective} {noun}" gorm:"not null"`
	CreatedAt time.Time `fake:"skip" gorm:"not null"`
	UpdatedAt time.Time `fake:"skip" gotm:"not null"`
}

// FakeCabinet returns a cabinet with fake data which can be inserted into the database.
func FakeCabinet() (*Cabinet, error) {
	var r Cabinet
	err := gofakeit.Struct(&r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
