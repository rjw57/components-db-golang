package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) CabinetsList(ctx *gin.Context, params CabinetsListParams) {
	s := "cabinet-2"
	resp := CabinetList{
		Items: &[]CabinetSummary{
			{Id: &s},
		},
	}

	ctx.JSON(http.StatusOK, resp)
}

func (Server) CabinetGet(ctx *gin.Context, cabinetId string) {
	resp := CabinetDetail{
		Id:      &cabinetId,
		Drawers: &[]DrawerDetail{},
	}

	ctx.JSON(http.StatusOK, resp)
}
