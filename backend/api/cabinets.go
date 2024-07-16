package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/rs/zerolog/log"

	"github.com/rjw57/components-db-golang/backend/db"
	"github.com/rjw57/components-db-golang/backend/db/schema/components/public/table"
	"github.com/rjw57/components-db-golang/backend/middleware"
)

func (s Server) CabinetsList(ctx *gin.Context, params CabinetsListParams) {
	tx := middleware.Tx(ctx)

	pageSize := DefaultPageSize
	if params.Limit != nil {
		pageSize = *params.Limit
	}

	items := []CabinetSummary{}
	stmt := table.Cabinets.
		SELECT(
			table.Cabinets.UUID.AS("CabinetSummary.Id"),
			table.Cabinets.Name.AS("CabinetSummary.Name"),
		).
		LIMIT(int64(pageSize + 1))

	if params.Cursor != nil {
		stmt = CabinetsStartingAtUUID(stmt, *params.Cursor)
	}

	if err := stmt.QueryContext(ctx, tx, &items); err != nil {
		log.Error().Err(err).Msg("Error querying cabinets")
		ctx.AbortWithStatus(500)
		return
	}

	resp := CabinetList{}
	if len(items) > pageSize {
		resp.NextCursor = items[pageSize].Id
		items = items[:pageSize]
	}
	resp.Items = &items

	ctx.JSON(http.StatusOK, resp)
}

func (s Server) CabinetGet(ctx *gin.Context, cabinetId UUID) {
	tx := middleware.Tx(ctx)

	c, err := db.MakeAndInsertFakeCabinet(tx)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	log.Info().Any("cabinet", c).Msg("Made fake cabinet")

	stmt := table.Cabinets.
		SELECT(
			table.Cabinets.UUID.AS("CabinetDetail.Id"),
			table.Cabinets.Name.AS("CabinetDetail.Name"),
		).
		WHERE(table.Cabinets.ID.EQ(pg.Int(c.ID)))

	resp := CabinetDetail{}
	if err := stmt.QueryContext(ctx, tx, &resp); err != nil {
		log.Error().Err(err).Msg("Failed to query cabinet")
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
