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

func (p populationHandler) GetPopulationStatistics(g *gin.Context) {
	popData, err := p.service.GetPopulationStatistics()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Module":   "handler",
			"Function": "GetPopulationStatistics",
		})
		logrus.Error(err)
	}
	g.JSON(200, popData)
}
