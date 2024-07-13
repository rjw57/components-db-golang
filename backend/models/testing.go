package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// OpenTestingDatabase opens a new Gorm session onto a testing database. If the TESTING_DATABASE_DSN
// environment variable is set, that is used as the connection DSN otherwise an ephemeral database
// is created via docker.
//
// A new pointer to a new gorm.DB instance is returned along with a function which must be called to
// close the instance.
func OpenTestingDatabase() (*gorm.DB, func() error, error) {
	logger := logger.New(
		log.Default(),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)

	dsn, ok := os.LookupEnv("TESTING_DATABASE_DSN")
	if !ok {
		log.Print("TESTING_DATABASE_DSN not set. Attempting to start temporary database via docker.")
		return openTemporaryDatabase(logger)
	}

	log.Print("Opening testing database from dsn specified in TESTING_DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger})
	if err != nil {
		return nil, nil, err
	}

	return db, (func() error { return nil }), nil
}

func openTemporaryDatabase(logger logger.Interface) (*gorm.DB, func() error, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, fmt.Errorf("Error getting docker pool: %w", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("Error connecting to docker pool: %w", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "16",
		Env: []string{
			"POSTGRES_DB=testing",
			"POSTGRES_USER=testing-user",
			"POSTGRES_PASSWORD=testing-pass",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container is removed from the file system
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		return nil, nil, fmt.Errorf("Error creating database container: %w", err)
	}

	cleanup := func() error {
		log.Print("Purging test database container")
		if err := resource.Close(); err != nil {
			return err
		}
		return nil
	}

	var db *gorm.DB
	pool.MaxWait = 5 * time.Minute
	if err := pool.Retry(func() error {
		var err error
		dsn := fmt.Sprintf(
			"postgres://testing-user:testing-pass@%s/testing",
			resource.GetHostPort("5432/tcp"),
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, nil, fmt.Errorf("Error connecting to database container: %w", err)
	}

	return db, cleanup, nil
}

type ModelSuite struct {
	suite.Suite
	DB *gorm.DB
}

func (s *ModelSuite) SetupSuite() {
	db, dbClose, err := OpenTestingDatabase()
	if err != nil {
		s.T().Errorf("Error opening test database: %s", err)
		s.T().FailNow()
		return
	}
	s.T().Cleanup(func() {
		log.Print("Closing database connection")
		dbClose()
	})

	log.Print("Starting suite transaction")
	s.DB = db.Begin()
	if err := s.DB.SavePoint("testsuite").Error; err != nil {
		s.T().Errorf("Error starting transaction: %s", err)
		s.T().FailNow()
		return
	}

	log.Print("Applying migrations")
	if err := s.DB.AutoMigrate(&Cabinet{}); err != nil {
		s.T().Errorf("Error migrating database: %s", err)
		s.T().FailNow()
		return
	}
}

func (s *ModelSuite) TearDownSuite() {
	log.Print("Rolling back suite transaction")
	if err := s.DB.RollbackTo("testsuite").Error; err != nil {
		s.T().Errorf("Error rolling back transaction: %s", s.DB.Error)
		s.T().FailNow()
		return
	}
	s.DB = nil
}

func (s *ModelSuite) BeforeTest(suiteName, testName string) {
	log.Print("Starting nested test transaction")
	if err := s.DB.SavePoint("test").Error; err != nil {
		s.T().Errorf("Error starting transaction: %s", err)
		s.T().FailNow()
		return
	}
}

func (s *ModelSuite) AfterTest(suiteName, testName string) {
	log.Print("Rolling back nested test transaction")
	s.DB.RollbackTo("test")
}
