package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Server) StatusGet(ctx *gin.Context) {
	ok := "ok"
	resp := ServerStatus{Status: &ok}
	ctx.JSON(http.StatusOK, resp)
}
