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

func (p *populationHandler) generateLogger(err error, functionName string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"Module":  "Handler",
		"Funtion": functionName,
	})
}

func (p populationHandler) GetPopulationStatistics(g *gin.Context) {
	popData, err := p.service.GetPopulationStatistics()
	if err == nil {
		g.JSON(http.StatusOK, popData)
		p.generateLogger(err, "GetPopulationStatistics")
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		g.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})
		p.generateLogger(err, "GetPopulationStatistics")
		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})
	p.generateLogger(err, "GetPopulationStatistics")
}

func (p populationHandler) GetAllCandidateInfo(g *gin.Context) {
	popData, err := p.service.GetAllCandidateInfo()
	if err == nil {
		g.JSON(http.StatusOK, popData)
		p.generateLogger(err, "GetAllCandidateInfo")
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		g.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})
		p.generateLogger(err, "GetAllCandidateInfo")
		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})
	p.generateLogger(err, "GetAllCandidateInfo")
}
