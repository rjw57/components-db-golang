//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Cabinets = newCabinetsTable("public", "cabinets", "Cabinet")

type cabinetsTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	UUID      postgres.ColumnString
	Name      postgres.ColumnString
	CreatedAt postgres.ColumnTimestampz
	UpdatedAt postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CabinetsTable struct {
	cabinetsTable

	EXCLUDED cabinetsTable
}

// AS creates new CabinetsTable with assigned alias
func (a CabinetsTable) AS(alias string) *CabinetsTable {
	return newCabinetsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CabinetsTable with assigned schema name
func (a CabinetsTable) FromSchema(schemaName string) *CabinetsTable {
	return newCabinetsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CabinetsTable with assigned table prefix
func (a CabinetsTable) WithPrefix(prefix string) *CabinetsTable {
	return newCabinetsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CabinetsTable with assigned table suffix
func (a CabinetsTable) WithSuffix(suffix string) *CabinetsTable {
	return newCabinetsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCabinetsTable(schemaName, tableName, alias string) *CabinetsTable {
	return &CabinetsTable{
		cabinetsTable: newCabinetsTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newCabinetsTableImpl("", "excluded", ""),
	}
}

func newCabinetsTableImpl(schemaName, tableName, alias string) cabinetsTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		UUIDColumn      = postgres.StringColumn("uuid")
		NameColumn      = postgres.StringColumn("name")
		CreatedAtColumn = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn = postgres.TimestampzColumn("updated_at")
		allColumns      = postgres.ColumnList{IDColumn, UUIDColumn, NameColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns  = postgres.ColumnList{UUIDColumn, NameColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return cabinetsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		UUID:      UUIDColumn,
		Name:      NameColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
