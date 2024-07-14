package model

import (
	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

// FakeCabinet writes fake data into a Cabinet struct.
func FakeCabinet(c *Cabinet) error {
	return gofakeit.Struct(c)
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
func MakeAndInsertFakeCabinet(db *gorm.DB) (*Cabinet, error) {
	c, err := MakeFakeCabinet()
	if err != nil {
		return nil, err
	}
	if err := db.Create(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// FakeCabinets writes fake data into the elements of a slice of Cabinet structs.
func FakeCabinets(cs []Cabinet) error {
	for _, c := range cs {
		if err := FakeCabinet(&c); err != nil {
			return err
		}
	}
	return nil
}

// MakeFakeCabinets returns a slice of newly created fake Cabinet structs.
func MakeFakeCabinets(n uint) ([]Cabinet, error) {
	cs := make([]Cabinet, n)
	if err := FakeCabinets(cs); err != nil {
		return nil, err
	}
	return cs, nil
}

// MakeAndInsertFakeCabinets returns a slice of newly created fake Cabinet structs which have been
// inserted into the database.
func MakeAndInsertFakeCabinets(db *gorm.DB, n uint) ([]Cabinet, error) {
	cs, err := MakeFakeCabinets(n)
	if err != nil {
		return nil, err
	}
	if err = db.Create(cs).Error; err != nil {
		return nil, err
	}
	return cs, nil
}
