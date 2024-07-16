package api

import (
	"testing"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stretchr/testify/suite"

	"github.com/rjw57/components-db-golang/backend/db"
	. "github.com/rjw57/components-db-golang/backend/db/schema/components/public/table"
)

type ModelSuite struct{ db.ModelSuite }

func TestModelSuite(t *testing.T) {
	suite.Run(t, &ModelSuite{})
}

func (s *ModelSuite) TestCabinetSummaryFields() {
	c, err := db.MakeAndInsertFakeCabinet(s.Tx)
	s.Require().NoError(err, "Error inserting fake Cabinets")

	cResult := CabinetSummary{}
	t := Cabinets.AS("CabinetSummary")
	s.Require().NoError(
		t.
			SELECT(
				t.UUID.AS("CabinetSummary.Id"),
				t.Name,
			).
			WHERE(t.ID.EQ(postgres.Int(c.ID))).
			Query(s.Tx, &cResult),
		"Failed to query cabinets",
	)
	s.EqualValues(&c.UUID, cResult.Id, "UUID not populated")
	s.EqualValues(&c.Name, cResult.Name, "Name is incorrect")
}
