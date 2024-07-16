package api

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/rjw57/components-db-golang/backend/db"
	"github.com/rjw57/components-db-golang/backend/db/schema/components/public/model"
	"github.com/rjw57/components-db-golang/backend/db/schema/components/public/table"
)

type QueriesSuite struct{ db.ModelSuite }

func TestQueriesSuite(t *testing.T) {
	suite.Run(t, &QueriesSuite{})
}

func (s *QueriesSuite) TestStartingAtUUID() {
	cs, err := db.MakeAndInsertFakeCabinets(s.Tx, 100)
	s.Require().NoError(err, "Error creating fake cabinets")

	testStartingAtIdx := func(idx int) {
		var found_cs []model.Cabinet
		starting_c := cs[idx]
		stmt := CabinetsStartingAtUUID(
			table.Cabinets.SELECT(table.Cabinets.AllColumns),
			starting_c.UUID,
		)
		s.Require().NoError(stmt.Query(s.Tx, &found_cs), "Error fetching cabinets")
		s.EqualValues(len(cs)-idx, len(found_cs), "Incorrect number of results")
		for _, c := range found_cs {
			s.GreaterOrEqual(c.ID, starting_c.ID, "Cabinet should not have been selected")
		}
	}

	s.Run("partial results", func() { testStartingAtIdx(40) })
	s.Run("all results", func() { testStartingAtIdx(0) })
	s.Run("no results", func() { testStartingAtIdx(len(cs) - 1) })
	s.Run("non existent UUID returns no results", func() {
		var found_cs []model.Cabinet
		uuid, err := uuid.NewRandom()
		s.Require().NoError(err, "Failed to create UUID")
		stmt := CabinetsStartingAtUUID(
			table.Cabinets.SELECT(table.Cabinets.AllColumns),
			uuid,
		)
		s.Require().NoError(stmt.Query(s.Tx, &found_cs), "Error fetching cabinets")
		s.EqualValues(0, len(found_cs), "Incorrect number of results")
	})
}
