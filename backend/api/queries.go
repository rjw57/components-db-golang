package api

import (
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"github.com/rjw57/components-db-golang/backend/db/schema/components/public/table"
)

func StartingAtUUID(s pg.SelectStatement, t pg.ReadableTable, idCol pg.ColumnInteger, uuidCol pg.ColumnString, uuid uuid.UUID) pg.SelectStatement {
	sq := pg.SELECT(idCol).FROM(t).WHERE(uuidCol.EQ(pg.UUID(uuid))).LIMIT(1)
	return s.WHERE(idCol.GT_EQ(pg.IntExp(sq))).ORDER_BY(idCol.ASC())
}

func CabinetsStartingAtUUID(s pg.SelectStatement, uuid uuid.UUID) pg.SelectStatement {
	t := table.Cabinets
	return StartingAtUUID(s, t, t.ID, t.UUID, uuid)
}
