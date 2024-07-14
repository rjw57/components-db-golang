package model

import (
	"log"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/rjw57/components-db-golang/backend/test"
)

type ModelSuite struct {
	suite.Suite
	DB *gorm.DB
}

func (s *ModelSuite) SetupSuite() {
	db, dbClose, err := test.OpenTestingDatabase()
	if err != nil {
		s.T().Errorf("Error opening test database: %s", err)
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
