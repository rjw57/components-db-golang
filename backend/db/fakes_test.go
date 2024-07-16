package db

import (
	"testing"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/rjw57/components-db-golang/backend/db/schema/components/public/table"
)

type FakesSuite struct{ ModelSuite }

func TestFakesSuite(t *testing.T) {
	suite.Run(t, &FakesSuite{})
}

func (s *FakesSuite) RequireCabinetCount(c int, msgAndArgs ...interface{}) {
	r := struct{ Count int64 }{}
	s.Require().NoError(table.Cabinets.
		SELECT(postgres.COUNT(table.Cabinets.ID).AS("count")).
		Query(s.Tx, &r),
	)
	s.Require().EqualValues(c, r.Count, msgAndArgs...)
}

func (s *FakesSuite) TestMakeFakeCabinet() {
	c, err := MakeFakeCabinet()
	s.Require().NoError(err, "Error creating fake cabinet")
	s.EqualValues(uuid.UUID{}, c.UUID, "Cabinet UUID is not zero value")
	s.EqualValues(c.ID, 0, "Cabinet id is not zero value")
	s.Equal(time.Time{}, c.CreatedAt, "Cabinet created at is not zero value")
	s.Equal(time.Time{}, c.UpdatedAt, "Cabinet updated at is not zero value")
	s.NotEqual("", c.Name, "Cabinet name is blank")
}

func (s *FakesSuite) TestMakeAndInsertFakeCabinet() {
	s.RequireCabinetCount(0, "Cabinets present in database")
	c, err := MakeAndInsertFakeCabinet(s.Tx)
	s.Require().NoError(err, "Error inserting fake Cabinet")
	s.NotEqual(uuid.UUID{}, c.UUID, "Cabinet UUID was not set")
	s.NotEqualValues(0, c.ID, "Cabinet ID was not set")
	s.RequireCabinetCount(1, "Cabinet not inserted in database")
}

func (s *FakesSuite) TestMakeFakeCabinets() {
	const n int = 10
	cs, err := MakeFakeCabinets(n)
	s.Require().NoError(err, "Error creating fake cabinets")
	for _, c := range cs {
		s.EqualValues(uuid.UUID{}, c.UUID, "Cabinet UUID is not zero value")
		s.EqualValues(c.ID, 0, "Cabinet id is not zero value")
		s.Equal(time.Time{}, c.CreatedAt, "Cabinet created at is not zero value")
		s.Equal(time.Time{}, c.UpdatedAt, "Cabinet updated at is not zero value")
		s.NotEqual("", c.Name, "Cabinet name is blank")
	}
}

func (s *FakesSuite) TestMakeAndInsertFakeCabinets() {
	const n int = 10
	s.RequireCabinetCount(0, "Cabinets present in database")
	cs, err := MakeAndInsertFakeCabinets(s.Tx, n)
	s.Require().NoError(err, "Error inserting fake Cabinet")
	s.EqualValues(n, len(cs), "Incorrect number of Cabinets created")
	for _, c := range cs {
		s.NotEqual(uuid.UUID{}, c.UUID, "Cabinet UUID was not set")
		s.NotEqualValues(0, c.ID, "Cabinet ID was not set")
		s.NotEqual("", c.Name, "Cabinet name is blank")
	}
	s.RequireCabinetCount(len(cs), "Incorrect number of Cabinets inserted")
}
