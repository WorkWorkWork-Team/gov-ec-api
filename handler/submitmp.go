package handler

import (
	"net/http"

	model "github.com/WorkWorkWork-Team/gov-ec-api/models"
	"github.com/WorkWorkWork-Team/gov-ec-api/service"
	"github.com/gin-gonic/gin"
)

type submitmpHandler struct {
	submitmpService service.SubmitmpService
}

func NewSubmitMpHandler(submitmpService service.SubmitmpService) submitmpHandler {
	return submitmpHandler{
		submitmpService: submitmpService,
	}
}

var mp model.SubmitMp

func (a *submitmpHandler) SubmitMp(g *gin.Context) {
	g.BindJSON(&mp)
	err := a.submitmpService.SubmitMp(mp.CitizenID)
	if err != nil {
		g.Status(http.StatusInternalServerError)
		return
	}
	g.Status(http.StatusOK)
	return
}
