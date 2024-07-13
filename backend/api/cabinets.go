package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/rjw57/components-db-golang/backend/models"
)

func (s Server) CabinetsList(ctx *gin.Context, params CabinetsListParams) {
	db := s.DB

	pageSize := DefaultPageSize
	if params.Limit != nil {
		pageSize = *params.Limit
	}
	db = db.Limit(pageSize)

	if params.Cursor != nil {
		db = db.Scopes(models.StartingAtUUID(*params.Cursor))
	}

	var items []CabinetSummary
	result := db.Find(&items)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Error querying cabinets")
		ctx.AbortWithError(500, result.Error)
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
	c := &models.Cabinet{}
	err := models.FakeCabinet(c)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	log.Info().Any("cabinet", c).Msg("Made fake cabinet")

	r := s.DB.Create(c)
	if r.Error != nil {
		ctx.AbortWithError(500, r.Error)
		return
	}

	log.Info().Any("cabinet", c).Msg("Inserted cabinet")
	resp := CabinetDetail{
		Id:      &c.UUID,
		Drawers: &[]DrawerDetail{},
		Name:    &c.Name,
	}

	ctx.JSON(http.StatusOK, resp)
}
