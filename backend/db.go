package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// OpenDatabase opens a new Gorm session based on the environment.
func OpenDatabase() (*gorm.DB, error) {
	dsn, ok := os.LookupEnv("DATABASE_DSN")
	if !ok {
		return nil, fmt.Errorf("DATABASE_DSN environment variable not set")
	}

	// TODO: disable in production?
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger})
}
