//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=oapi-codegen.yaml ../../openapi/spec.yaml
package api

const DefaultPageSize = 100

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (CabinetSummary) TableName() string {
	return "cabinets"
}
