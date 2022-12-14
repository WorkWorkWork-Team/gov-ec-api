package handler

import (
	"fmt"
	"net/http"

	"github.com/WorkWorkWork-Team/gov-ec-api/model"
	"github.com/WorkWorkWork-Team/gov-ec-api/service"
	"github.com/gin-gonic/gin"
)

type SubmitmpHandler struct {
	submitmpService service.SubmitmpService
}

func NewSubmitMpHandler(submitmpService service.SubmitmpService) *SubmitmpHandler {
	return &SubmitmpHandler{
		submitmpService: submitmpService,
	}
}

func (a *SubmitmpHandler) SubmitMp(g *gin.Context) {
	var mp model.SubmitMp

	err := g.BindJSON(&mp)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body.",
		})
		return
	}
	err = a.submitmpService.SubmitMp(mp.CitizenID)
	if err != nil {
		if err == service.ErrCitizenIDNotFound {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprint("Internal Server Error: ", err),
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
