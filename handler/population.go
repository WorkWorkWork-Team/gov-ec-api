package handler

import (
	"database/sql"
	"errors"
	"net/http"

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

func (p *populationHandler) generateLogger(functionName string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"Module":  "Handler",
		"Funtion": functionName,
	})
}

func (p populationHandler) GetPopulationStatistics(g *gin.Context) {
	popData, err := p.service.GetPopulationStatistics()
	if err == nil {
		g.JSON(http.StatusOK, popData)
		p.generateLogger("GetPopulationStatistics").Info("return population statistics")
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		g.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})
		p.generateLogger("GetPopulationStatistics").Error(err)
		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})
	p.generateLogger("GetPopulationStatistics").Error(err)
}

func (p populationHandler) GetAllCandidateInfo(g *gin.Context) {
	popData, err := p.service.GetAllCandidateInfo()
	if err == nil {
		g.JSON(http.StatusOK, popData)
		p.generateLogger("GetAllCandidateInfo").Info("return all candidate info")
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		g.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})
		p.generateLogger("GetAllCandidateInfo").Error(err)
		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})
	p.generateLogger("GetAllCandidateInfo").Error(err)
}
