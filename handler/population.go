package handler

import (
	"net/http"

	"github.com/WorkWorkWork-Team/gov-ec-api/service"
	"github.com/gin-gonic/gin"
)

type populationHandler struct {
	service service.PopulationService
}

func NewPopulationHandler(populationService service.PopulationService) populationHandler {
	return populationHandler{
		service: populationService,
	}
}

func (p populationHandler) GetPopulationStatistics(g *gin.Context) {
	popData := p.service.GetTotalPopulation()
	g.JSON(http.StatusOK, gin.H{
		"test": popData,
	})
}
