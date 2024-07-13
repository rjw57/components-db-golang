package models

import (
	"github.com/google/uuid"
)

func (s *ModelSuite) TestCabinetAutoUUID() {
	c := &Cabinet{}
	if s.NoError(FakeCabinet(c), "Error creating fake cabinet") {
		s.Equal(c.UUID, uuid.UUID{}, "Cabinet UUID is not zero value")
		/*
			if s.NoError(s.db.Create(c).Error, "Error inserting cabinet") {
				s.NotEqual(c.UUID, uuid.UUID{}, "Cabinet UUID was not updated")
			}
		*/
	}
}
