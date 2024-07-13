package api

import (
	"testing"

	"github.com/rjw57/components-db-golang/backend/models"
	"github.com/stretchr/testify/suite"
)

type ModelSuite struct{ models.ModelSuite }

func TestModelSuite(t *testing.T) {
	suite.Run(t, &ModelSuite{})
}

func (s *ModelSuite) TestCabinetSummaryFields() {
	c, err := models.MakeAndInsertFakeCabinet(s.DB)
	s.Require().NoError(err, "Error inserting fake Cabinets")

	cs := &CabinetSummary{}
	s.Require().NoError(s.DB.Where("id = ?", c.ID).Take(cs).Error, "Error querying database")
	s.EqualValues(&c.UUID, cs.Id, "UUID not populated")
	s.Greater(len(*cs.Name), 0, "Name is empty")
}
