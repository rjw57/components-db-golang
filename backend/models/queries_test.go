package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type QueriesSuite struct{ ModelSuite }

func TestQueriesSuite(t *testing.T) {
	suite.Run(t, &QueriesSuite{})
}

func (s *QueriesSuite) TestStartingAtUUID() {
	cs, err := MakeAndInsertFakeCabinets(s.DB, 100)
	if !s.NoError(err, "Error creating fake cabinets") {
		return
	}

	testStartingAtIdx := func(idx int) {
		var found_cs []Cabinet
		starting_c := cs[idx]
		err = s.DB.Scopes(StartingAtUUID(starting_c.UUID)).Find(&found_cs).Error
		if !s.NoError(err, "Error fetching cabinets") {
			return
		}

		s.EqualValues(len(cs)-idx, len(found_cs), "Incorrect number of results")
		for _, c := range found_cs {
			s.LessOrEqual(starting_c.ID, c.ID, "Cabinet should not have been selected")
		}
	}

	s.Run("partial results", func() { testStartingAtIdx(40) })
	s.Run("all results", func() { testStartingAtIdx(0) })
	s.Run("no results", func() { testStartingAtIdx(len(cs) - 1) })
	s.Run("non existent UUID returns no results", func() {
		var found_cs []Cabinet
		uuid, err := uuid.NewRandom()
		if !s.NoError(err, "Failed to create UUID") {
			return
		}
		err = s.DB.Scopes(StartingAtUUID(uuid)).Find(&found_cs).Error
		if !s.NoError(err, "Error fetching cabinets") {
			return
		}

		s.EqualValues(0, len(found_cs), "Incorrect number of results")
	})
}
