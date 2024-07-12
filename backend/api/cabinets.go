package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/rjw57/components-db-golang/backend/models"
)

func (s Server) CabinetsList(ctx *gin.Context, params CabinetsListParams) {
	var cabinets []models.Cabinet

	tx := s.DB.Limit(PageSize + 1).Order("id ASC")
	if params.Cursor != nil {
		c, err := uuid.Parse(*params.Cursor)
		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}
		tx = tx.Where("id >= (?)", s.DB.Table("cabinets").Select("id").Where("uuid = ?", c))
	}

	result := tx.Find(&cabinets)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Error querying cabinets")
		ctx.AbortWithError(500, result.Error)
		return
	}

	items := make([]CabinetSummary, 0, PageSize)
	for i, cab := range cabinets {
		if i >= PageSize {
			break
		}
		id := cab.UUID.String()
		items = append(items, CabinetSummary{Id: &id})
	}

	resp := CabinetList{Items: &items}

	if len(cabinets) > PageSize {
		var s = cabinets[PageSize].UUID.String()
		resp.NextCursor = &s
	}

	ctx.JSON(http.StatusOK, resp)
}

func (s Server) CabinetGet(ctx *gin.Context, cabinetId string) {
	resp := CabinetDetail{
		Id:      &cabinetId,
		Drawers: &[]DrawerDetail{},
	}

	ctx.JSON(http.StatusOK, resp)
}
