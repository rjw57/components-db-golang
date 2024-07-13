package models

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type FakesSuite struct{ ModelSuite }

func TestFakesSuite(t *testing.T) {
	suite.Run(t, &FakesSuite{})
}

func (s *FakesSuite) TestMakeFakeCabinet() {
	c, err := MakeFakeCabinet()
	s.Require().NoError(err, "Error creating fake cabinet")
	s.Equal(c.UUID, uuid.UUID{}, "Cabinet UUID is not zero value")
	s.EqualValues(c.ID, 0, "Cabinet id is not zero value")
	s.Equal(time.Time{}, c.CreatedAt, "Cabinet created at is not zero value")
	s.Equal(
		time.Time{},
		c.UpdatedAt,
		"Cabinet updated at is not zero value",
	)
	s.NotEqual("", c.Name, "Cabinet name is blank")
}

func (s *FakesSuite) TestMakeAndInsertFakeCabinet() {
	var count int64
	s.Require().NoError(s.DB.Model(&Cabinet{}).Count(&count).Error, "Error counting cabinets")
	s.EqualValues(count, 0, "Cabinets exist in database")

	c, err := MakeAndInsertFakeCabinet(s.DB)
	s.Require().NoError(err, "Error inserting fake Cabinet")
	s.NotEqual(uuid.UUID{}, c.UUID, "Cabinet UUID was not set")
	s.NotEqualValues(0, c.ID, "Cabinet ID was not set")

	s.Require().NoError(s.DB.Model(&Cabinet{}).Count(&count).Error, "Error counting cabinets")
	s.EqualValues(1, count, "Not exactly one Cabinet in database")
}

func (s *FakesSuite) TestMakeAndInsertFakeCabinets() {
	var n uint = 100

	var count int64
	s.Require().NoError(s.DB.Model(&Cabinet{}).Count(&count).Error, "Error counting cabinets")
	s.EqualValues(0, count, "Cabinets exist in database")

	cs, err := MakeAndInsertFakeCabinets(s.DB, n)
	s.Require().NoError(err, "Error inserting fake Cabinet")
	s.EqualValues(n, len(cs), "Incorrect number of Cabinets created")

	s.Require().NoError(s.DB.Model(&Cabinet{}).Count(&count).Error, "Error counting cabinets")
	s.EqualValues(n, count, "Incorrect number of Cabinets inserted")
}
