package handler

import (
	"github.com/WorkWorkWork-Team/gov-ec-api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type populationHandler struct {
	service service.PopulationService
}

func NewPopulationHandler(populationService service.PopulationService) populationHandler {
	return populationHandler{
		service: populationService,
	}
}

func (p *populationHandler) errorMessage(err error, functionName string) {
	logrus.WithFields(logrus.Fields{
		"Module":  "Handler",
		"Funtion": functionName,
	}).Error(err)
}

func (p populationHandler) GetPopulationStatistics(g *gin.Context) {
	popData, err := p.service.GetPopulationStatistics()
	if err != nil {
		p.errorMessage(err, "GetPopulationStatistics")
	}
	g.JSON(200, popData)
}

func (p populationHandler) GetAllCandidateInfo(g *gin.Context) {
	popData, err := p.service.GetAllCandidateInfo()
	if err != nil {
		p.errorMessage(err, "GetAllCandidateInfo")
	}
	g.JSON(200, popData)
}
