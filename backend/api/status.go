package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Server) StatusGet(ctx *gin.Context) {
	resp := ServerStatus{}
	ctx.JSON(http.StatusOK, resp)
}
