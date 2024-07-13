package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StartingAtUUID returns a Gorm scope function which queries for instances of a model whose numeric
// id is greater than or equal to the model with the matching UUID. If there is no matching UUID, no
// results are returned.
func StartingAtUUID(uuid uuid.UUID) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		subquery := db.Session(&gorm.Session{}).Select("id").Where("uuid = ?", uuid).Limit(1)
		return db.Where("id >= (?)", subquery).Order("id ASC")
	}
}
