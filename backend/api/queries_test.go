package api

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/rjw57/components-db-golang/backend/model"
)

type QueriesSuite struct{ ModelSuite }

func TestQueriesSuite(t *testing.T) {
	suite.Run(t, &QueriesSuite{})
}

func (s *QueriesSuite) TestStartingAtUUID() {
	cs, err := model.MakeAndInsertFakeCabinets(s.DB, 100)
	s.Require().NoError(err, "Error creating fake cabinets")

	testStartingAtIdx := func(idx int) {
		var found_cs []model.Cabinet
		starting_c := cs[idx]
		s.Require().NoError(
			s.DB.Scopes(StartingAtUUID(starting_c.UUID)).Find(&found_cs).Error,
			"Error fetching cabinets",
		)

		s.EqualValues(len(cs)-idx, len(found_cs), "Incorrect number of results")
		for _, c := range found_cs {
			s.LessOrEqual(starting_c.ID, c.ID, "Cabinet should not have been selected")
		}
	}

	s.Run("partial results", func() { testStartingAtIdx(40) })
	s.Run("all results", func() { testStartingAtIdx(0) })
	s.Run("no results", func() { testStartingAtIdx(len(cs) - 1) })
	s.Run("non existent UUID returns no results", func() {
		var found_cs []model.Cabinet
		uuid, err := uuid.NewRandom()
		s.Require().NoError(err, "Failed to create UUID")
		err = s.DB.Scopes(StartingAtUUID(uuid)).Find(&found_cs).Error
		s.Require().NoError(err, "Error fetching cabinets")
		s.EqualValues(0, len(found_cs), "Incorrect number of results")
	})
}
