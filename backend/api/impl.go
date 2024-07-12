package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) CabinetsList(ctx *gin.Context) {
	s := "cabinet-2"
	resp := CabinetList{
		Items: &[]CabinetSummary{
			{Id: &s},
		},
	}

	ctx.JSON(http.StatusOK, resp)
}
