package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// OpenDatabase opens a new Gorm session based on the environment.
func OpenDatabase() (*sql.DB, error) {
	dsn, ok := os.LookupEnv("DATABASE_DSN")
	if !ok {
		return nil, fmt.Errorf("DATABASE_DSN environment variable not set")
	}
	return sql.Open("postgres", dsn)
}
