package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type CabinetSuite struct{ ModelSuite }

func TestCabinetSuite(t *testing.T) {
	suite.Run(t, &CabinetSuite{})
}

func (s *CabinetSuite) TestCabinetAutoUUID() {
	c := &Cabinet{}
	if !s.NoError(FakeCabinet(c), "Error creating fake cabinet") {
		return
	}
	s.Equal(c.UUID, uuid.UUID{}, "Cabinet UUID is not zero value")
	if !s.NoError(s.DB.Create(c).Error, "Error inserting cabinet") {
		return
	}
	s.NotEqual(c.UUID, uuid.UUID{}, "Cabinet UUID was not updated")
}
