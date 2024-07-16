package db

import (
	"github.com/rjw57/components-db-golang/backend/db/schema/components/public/table"
)

var CabinetInsertableColumns = table.Cabinets.AllColumns.Except(
	table.Cabinets.ID,
	table.Cabinets.UUID,
	table.Cabinets.CreatedAt,
	table.Cabinets.UpdatedAt,
)
