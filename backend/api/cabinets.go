package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/rjw57/components-db-golang/backend/models"
)

func (s Server) CabinetsList(ctx *gin.Context, params CabinetsListParams) {
	var cabinets []models.Cabinet

	pageSize := DefaultPageSize
	if params.Limit != nil {
		pageSize = *params.Limit
	}

	tx := s.DB.Limit(pageSize + 1).Order("id ASC")
	if params.Cursor != nil {
		tx = tx.Where("id >= (?)", s.DB.Table("cabinets").Select("id").Where("uuid = ?", params.Cursor))
	}

	result := tx.Find(&cabinets)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Error querying cabinets")
		ctx.AbortWithError(500, result.Error)
		return
	}

	items := make([]CabinetSummary, 0, pageSize)
	for i, cab := range cabinets {
		if i >= pageSize {
			break
		}
		items = append(items, CabinetSummary{Id: &cab.UUID})
	}

	resp := CabinetList{Items: &items}

	if len(cabinets) > pageSize {
		resp.NextCursor = &cabinets[pageSize].UUID
	}

	ctx.JSON(http.StatusOK, resp)
}

func (s Server) CabinetGet(ctx *gin.Context, cabinetId UUID) {
	c, err := models.FakeCabinet()
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
