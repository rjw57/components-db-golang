package db

import (
	"database/sql"
	"log"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stretchr/testify/suite"

	"github.com/rjw57/components-db-golang/backend/db/schema/components/public/table"
	"github.com/rjw57/components-db-golang/backend/test"
)

type ModelSuite struct {
	suite.Suite
	Db *sql.DB
	Tx *sql.Tx
}

func (s *ModelSuite) SetupSuite() {
	db, dbClose, err := test.OpenTestingDatabase()
	if err != nil {
		s.T().Errorf("Error opening test database: %s", err)
		s.T().FailNow()
	}
	s.T().Cleanup(func() {
		log.Print("Closing database connection")
		dbClose()
	})
	s.Db = db
}

func (s *ModelSuite) BeforeTest(suiteName, testName string) {
	var err error

	log.Print("Starting test transaction")
	if tx, err := s.Db.Begin(); err != nil {
		s.T().Errorf("Error starting transaction: %s", err)
		s.T().FailNow()
	} else {
		s.Tx = tx
	}

	log.Print("Dropping any existing data in the test database.")
	_, err = table.Cabinets.DELETE().WHERE(postgres.Bool(true)).Exec(s.Tx)
	s.Require().NoError(err)
}

func (s *ModelSuite) AfterTest(suiteName, testName string) {
	log.Print("Rolling back test transaction")
	if err := s.Tx.Rollback(); err != nil {
		s.T().Errorf("Error rolling transaction back: %s", err)
	}
	s.Tx = nil
}
