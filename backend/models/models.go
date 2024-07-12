package models

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

type Cabinet struct {
	ID        uint      `fake:"skip"`
	UUID      uuid.UUID `fake:"skip" gorm:"index;type:uuid;default:gen_random_uuid()"`
	Name      string    `fake:"{adjective} {noun}"`
	CreatedAt time.Time `fake:"skip"`
	UpdatedAt time.Time `fake:"skip"`
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
