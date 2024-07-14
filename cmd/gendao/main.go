package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:           "backend/query",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	// Initialize a *gorm.DB instance
	db, err := OpenDatabase()
	if err != nil {
		panic(err)
	}

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("cabinets"),
	)

	// Execute the generator
	g.Execute()
}

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
