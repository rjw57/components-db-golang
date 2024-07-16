package db

import (
	"errors"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"

	. "github.com/rjw57/components-db-golang/backend/db/schema/components/public/model"
	"github.com/rjw57/components-db-golang/backend/db/schema/components/public/table"
)

// insertCabinets() returns an insert statement which inserts the non-generated fields for a
// Cabinet and returns the inserted Cabinet instances.
func insertCabinets() postgres.InsertStatement {
	return table.Cabinets.
		INSERT(CabinetInsertableColumns).
		RETURNING(table.Cabinets.AllColumns)
}

// FakeCabinet writes fake data into a Cabinet struct.
func FakeCabinet(c *Cabinet) error {
	if c == nil {
		return errors.New("nil pointer passed")
	}
	c.Name = fmt.Sprintf("%s %s", gofakeit.Adjective(), gofakeit.Name())
	return nil
}

// MakeFakeCabinet returns a new Cabinet with fake data.
func MakeFakeCabinet() (*Cabinet, error) {
	c := &Cabinet{}
	if err := FakeCabinet(c); err != nil {
		return nil, err
	}
	return c, nil
}

// MakeAndInsertFakeCabinet returns a new Cabinet with fake data which has been inserted into the
// database.
func MakeAndInsertFakeCabinet(db qrm.DB) (*Cabinet, error) {
	c, err := MakeFakeCabinet()
	if err != nil {
		return nil, err
	}
	if err := insertCabinets().MODEL(c).Query(db, c); err != nil {
		return nil, err
	}
	return c, nil
}

// FakeCabinets writes fake data into the elements of a slice of Cabinet structs.
func FakeCabinets(cs []Cabinet) error {
	for i := range cs {
		if err := FakeCabinet(&cs[i]); err != nil {
			return err
		}
	}
	return nil
}

// MakeFakeCabinets returns a slice of newly created fake Cabinet structs.
func MakeFakeCabinets(n int) ([]Cabinet, error) {
	cs := make([]Cabinet, n)
	if err := FakeCabinets(cs); err != nil {
		return nil, err
	}
	return cs, nil
}

// MakeAndInsertFakeCabinets returns a slice of newly created fake Cabinet structs which have been
// inserted into the database.
func MakeAndInsertFakeCabinets(db qrm.DB, n int) ([]Cabinet, error) {
	cs, err := MakeFakeCabinets(n)
	if err != nil {
		return nil, err
	}
	insertedCs := []Cabinet{}
	if err := insertCabinets().MODELS(&cs).Query(db, &insertedCs); err != nil {
		return nil, err
	}
	return insertedCs, nil
}
