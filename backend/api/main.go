//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-codegen.yaml ../../openapi/spec.yaml
package api

import (
	"gorm.io/gorm"
)

const DefaultPageSize = 100

type Server struct {
	DB *gorm.DB
}

func NewServer(db *gorm.DB) Server {
	return Server{DB: db}
}

func (CabinetSummary) TableName() string {
	return "cabinets"
}
